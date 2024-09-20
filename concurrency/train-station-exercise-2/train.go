package main

import (
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
	// Passengers that did not complete the journey yet (on the train or waiting for it...)
	ActivePassengers []*Passenger
	// Total number of stops
	TotalStops int

	// Direction of the train
	GoingForward bool

	// Channels for boarding and disembarking passengers
	BoardChannel     chan *Passenger
	DisembarkChannel chan *Passenger
	IgnoreChannel    chan *Passenger

	CurrentStopChannel chan int

	// channel to stop the train
	StopChannel chan bool
}

func NewTrain(totalStops int, goingForward bool) *Train {
	return &Train{
		PassengersOnboard:  make([]*Passenger, 0),
		CurrentStopChannel: make(chan int),
		TotalStops:         totalStops,
		GoingForward:       goingForward,
		BoardChannel:       make(chan *Passenger),
		DisembarkChannel:   make(chan *Passenger),
		IgnoreChannel:      make(chan *Passenger),
		StopChannel:        make(chan bool),
	}
}

func (t *Train) BoardPassenger(p *Passenger) {
	t.PassengersOnboard = append(t.PassengersOnboard, p)
}

func (t *Train) DisembarkPassenger(p *Passenger) {
	for i, passenger := range t.PassengersOnboard {
		if passenger == p {
			t.PassengersOnboard = append(t.PassengersOnboard[:i], t.PassengersOnboard[i+1:]...)
			t.ActivePassengers = append(t.ActivePassengers[:i], t.ActivePassengers[i+1:]...)
			break
		}
	}
}

func (t *Train) addActivePassenger(p *Passenger) {
	t.ActivePassengers = append(t.ActivePassengers, p)
}

func shouldBoardPassenger(t *Train, currentStop int, p *Passenger) bool {
	if t.GoingForward {
		return p.Ticket.Origin == currentStop && p.Ticket.Origin < p.Ticket.Destination
	}

	return p.Ticket.Origin == currentStop && p.Ticket.Origin > p.Ticket.Destination
}

func shouldDisembarkPassenger(t *Train, currentStop int, p *Passenger) bool {
	if t.GoingForward {
		return p.Ticket.Destination == currentStop && p.Ticket.Origin < p.Ticket.Destination
	}

	return p.Ticket.Destination == currentStop && p.Ticket.Origin > p.Ticket.Destination
}

func (t *Train) stopTrain() {
	close(t.BoardChannel)
	close(t.DisembarkChannel)
	close(t.StopChannel)
	color.Red("-- No more passengers are awaiting, the train will stop... --")
}

func (t *Train) runTrain() {
	color.Cyan("[TRAIN MOVING] - The train is starting its journey...\n")
	currentStation := startingStop
	t.CurrentStopChannel <- currentStation

	for {
		// Simulate the train moving
		time.Sleep(2 * time.Second)

		color.Cyan("[TRAIN STOPPING] - Train in station %d with %d passengers! **\n", currentStation, len(t.PassengersOnboard))

		// Check if there are any passengers to board or disembark
		for i := 0; i < len(t.ActivePassengers); i++ {
			select {
			case p := <-t.BoardChannel:
				color.Green("[BOARDING] - %s is boarding! **\n", p.Name)
				t.BoardPassenger(p)
			case p := <-t.DisembarkChannel:
				color.Red("[DISEMBARKING] - %s is disembarking! **\n", p.Name)
				t.DisembarkPassenger(p)
			case p := <-t.IgnoreChannel:
				color.Yellow("[IGNORE] - %s Ignoring station! **\n", p.Name)
			case <-t.StopChannel:
				t.stopTrain()
				return
			default:
				// No action needed, continue to the next passenger
			}
		}

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

		color.Cyan("[TRAIN MOVING] - The train is moving. NEXT STATION: %d **\n", currentStation)
		t.CurrentStopChannel <- currentStation
	}
}
