package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fatih/color"
)

var totalStops = 5
var startingStop = 0
var numberOfPassengers = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	train := NewTrain(totalStops, true)

	broadcaster := NewTrainReportService(ctx, train.CurrentStopChannel)

	var wg sync.WaitGroup

	wg.Add(numberOfPassengers)

	for i := 0; i < numberOfPassengers; i++ {
		go func(passengerId int) {
			ticket := generatePassengerTicket()
			passenger := NewPassenger(fmt.Sprintf("Passenger-%d", passengerId), ticket)

			defer func() {
				color.Cyan("[PASSENGER] - %s has left the station!", passenger.Name)
				wg.Done()
			}()

			color.Cyan("[PASSENGER] - %s has a ticket from %d to %d!\n", passenger.Name, passenger.Ticket.Origin, passenger.Ticket.Destination)

			stationReportListener := broadcaster.Subscribe()
			train.addActivePassenger(passenger)

			for {
				select {
				case <-ctx.Done():
					// Exit if the context is done (e.g., due to cancellation)
					return
				case currentStation := <-stationReportListener:
					// Check if the passenger should board or disembark
					shouldBoard := shouldBoardPassenger(train, currentStation, passenger)
					shouldDisembark := shouldDisembarkPassenger(train, currentStation, passenger)

					if shouldBoard || shouldDisembark {
						if shouldBoard {
							// Send the passenger to the boarding channel
							train.BoardChannel <- passenger
						}

						if shouldDisembark {
							// Send the passenger to the disembarking channel
							train.DisembarkChannel <- passenger
							return
						}
					} else {
						// Passenger is ignoring the train report because it's not at the right station
						train.IgnoreChannel <- passenger
					}
				}
			}
		}(i)
	}

	go train.runTrain()

	wg.Wait()

	train.stopTrain()
}
