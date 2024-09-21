package main

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

// TrainService is an interface that defines the methods that a train service should implement
type TrainService interface {
	BoardPassenger(p *Passenger)
	DisembarkPassenger(p *Passenger)
}

type Train struct {
	// PassengersOnboard on the train
	PassengersOnboard []*Passenger
	// Active passengers on the train or stations
	ActivePassengers []*Passenger
	// Total number of stops
	TotalStops int

	// Direction of the train
	GoingForward bool

	// Channels for boarding and disembarking passengers
	PassengerResponseChannel chan *PassengerResponse

	// Channel to broadcast train position
	CurrentStopChannel chan int

	// channel to stop the train
	StopChannel chan bool
}

func NewTrain(totalStops int, goingForward bool) *Train {
	return &Train{
		PassengersOnboard:        make([]*Passenger, 0),
		TotalStops:               totalStops,
		GoingForward:             goingForward,
		PassengerResponseChannel: make(chan *PassengerResponse),
		CurrentStopChannel:       make(chan int),
		StopChannel:              make(chan bool),
	}
}

func (t *Train) BoardPassenger(p *Passenger) {
	t.PassengersOnboard = append(t.PassengersOnboard, p)
}

func (t *Train) DisembarkPassenger(p *Passenger) {
	for i, passenger := range t.PassengersOnboard {
		if passenger == p {
			t.PassengersOnboard = append(t.PassengersOnboard[:i], t.PassengersOnboard[i+1:]...)
			break
		}
	}
	t.RemoveActivePassenger(p)
}

func (t *Train) RemoveActivePassenger(p *Passenger) {
	for i, passenger := range t.ActivePassengers {
		if passenger == p {
			t.ActivePassengers = append(t.ActivePassengers[:i], t.ActivePassengers[i+1:]...)
			break
		}
	}
}

func shouldBoardPassenger(t *Train, currentStop int, p *Passenger) bool {
	if t.GoingForward {
		return p.Ticket.Origin == currentStop && p.Ticket.Origin < p.Ticket.Destination
	}

	return p.Ticket.Origin == currentStop && p.Ticket.Origin > p.Ticket.Destination
}

func (t *Train) passengerOnboard(p *Passenger) bool {
	for _, passenger := range t.PassengersOnboard {
		if passenger == p {
			return true
		}
	}

	return false
}

func shouldDisembarkPassenger(t *Train, currentStop int, p *Passenger) bool {
	passengerOnboard := t.passengerOnboard(p)
	if t.GoingForward {
		return p.Ticket.Destination == currentStop && p.Ticket.Origin < p.Ticket.Destination && passengerOnboard
	}

	return p.Ticket.Destination == currentStop && p.Ticket.Origin > p.Ticket.Destination && passengerOnboard
}

func (t *Train) stopTrain() {
	close(t.PassengerResponseChannel)
	close(t.CurrentStopChannel)
	close(t.StopChannel)
	color.Red("-- No more passengers are awaiting, the train will stop... --")
}

func (t *Train) runTrain() {
	color.Cyan("[TRAIN MOVING] - The train is starting its journey...\n")
	currentStation := startingStop

	for {
		// Simulate the train moving
		time.Sleep(1 * time.Second)

		t.CurrentStopChannel <- currentStation

		// Broadcast the current station to all passengers
		color.Cyan("[TRAIN STOPPING] - Train in station %d!\n", currentStation)

		var wg sync.WaitGroup
		passengerCount := len(t.ActivePassengers)

		// Add to the WaitGroup the number of passengers to expect responses from
		wg.Add(passengerCount)

		go func() {
			for i := 0; i < passengerCount; i++ {
				response := <-t.PassengerResponseChannel
				switch response.Action {
				case "board":
					t.BoardPassenger(response.Passenger)
					color.Green("[BOARDING] - Passenger %s boarded at station %d\n", response.Passenger.Name, currentStation)
				case "disembark":
					t.DisembarkPassenger(response.Passenger)
					color.Red("[DISEMBARKING] - Passenger %s disembarked at station %d\n", response.Passenger.Name, currentStation)
				case "ignore":
					color.Yellow("[IGNORING] - Passenger %s ignored the train at station %d\n", response.Passenger.Name, currentStation)
				}
				// Mark this response as handled
				wg.Done()
			}
		}()

		// Wait for all passenger responses before proceeding
		wg.Wait()

		// Move the train to the next stop
		if t.GoingForward {
			if currentStation < t.TotalStops {
				currentStation++
			} else {
				t.GoingForward = false
				currentStation--
			}
		} else {
			if currentStation > 0 {
				currentStation--
			} else {
				t.GoingForward = true
				currentStation++
			}
		}
	}
}
