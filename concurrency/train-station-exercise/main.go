package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

var totalStops = 5
var startingStop = 0

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create a train
	train := NewTrain(totalStops, true)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	// create passengers
	go func() {
		ticket := generatePassengerTicket()
		passenger := NewPassenger("Gonçalo", ticket)
		defer color.Cyan("[PASSENGER] - %s has left the station!", passenger.Name)
		defer wg.Done()

		color.Cyan("[PASSENGER] - %s has a ticket from %d to %d!\n", passenger.Name, passenger.Ticket.Origin, passenger.Ticket.Destination)

		for station := range train.CurrentStopChannel {
			if shouldBoardPassenger(train, station, passenger) {
				train.BoardChannel <- passenger
			}

			if shouldDisembarkPassenger(train, station, passenger) {
				train.DisembarkChannel <- passenger
				return
			}
		}
	}()

	// start the train
	go train.runTrain()

	wg.Wait()
	train.stopTrain()
}
