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
	// Passengers on the train
	Passengers []*Passenger
	// Total number of stops
	TotalStops int

	// Direction of the train
	GoingForward bool

	// Channels for boarding and disembarking passengers
	BoardChannel     chan *Passenger
	DisembarkChannel chan *Passenger

	CurrentStopChannel chan int

	// channel to stop the train
	StopChannel chan bool
}

func NewTrain(totalStops int, goingForward bool) *Train {
	return &Train{
		Passengers:         make([]*Passenger, 0),
		CurrentStopChannel: make(chan int),
		TotalStops:         totalStops,
		GoingForward:       goingForward,
		BoardChannel:       make(chan *Passenger),
		DisembarkChannel:   make(chan *Passenger),
		StopChannel:        make(chan bool),
	}
}

func (t *Train) BoardPassenger(p *Passenger) {
	t.Passengers = append(t.Passengers, p)
}

func (t *Train) DisembarkPassenger(p *Passenger) {
	for i, passenger := range t.Passengers {
		if passenger == p {
			t.Passengers = append(t.Passengers[:i], t.Passengers[i+1:]...)
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

		color.Cyan("[TRAIN STOPPING] - Train in station %d with %d passengers! **\n", currentStation, len(t.Passengers))

		// Check if there are any passengers to board or disembark
		select {
		case p := <-t.BoardChannel:
			color.Green("[BOARDING PASSENGER] - Passenger %s is boarding the train! **\n", p.Name)
			t.BoardPassenger(p)
		case p := <-t.DisembarkChannel:
			color.Red("[DISEMBARKING PASSENGER] - Passenger %s is disembarking the train! **\n", p.Name)
			t.DisembarkPassenger(p)
		case <-t.StopChannel:
			t.stopTrain()
		default:
		}

		// Move the train to the next stop
		if t.GoingForward {
			if currentStation < t.TotalStops {
				color.Cyan("[TRAIN MOVING FORWARD] - The train is moving. NEXT STATION: %d **\n", currentStation+1)
				currentStation++
			} else {
				color.Red("[END OF LINE] - The Train reached the final stop. NEXT STATION: %d **\n", currentStation-1)
				t.GoingForward = false
				currentStation--
			}
		} else {
			if currentStation > 0 {
				color.Cyan("[TRAIN MOVING BACK] - The train is moving. NEXT STATION: %d **\n", currentStation-1)
				currentStation--
			} else {
				color.Red("[END OF LINE] - The Train reached the final stop. NEXT STATION: %d **\n", currentStation+1)
				t.GoingForward = true
				currentStation++
			}
		}
		t.CurrentStopChannel <- currentStation
	}
}
