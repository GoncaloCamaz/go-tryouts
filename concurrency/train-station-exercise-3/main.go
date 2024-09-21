package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/fatih/color"
)

var totalStops = 5
var startingStop = 0
var numberOfPassengers = 10

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

			train.ActivePassengers = append(train.ActivePassengers, passenger)

			defer func() {
				color.Cyan("[PASSENGER] - %s has left the station!", passenger.Name)
				wg.Done()
			}()

			color.Cyan("[PASSENGER] - %s has a ticket from %d to %d!\n", passenger.Name, passenger.Ticket.Origin, passenger.Ticket.Destination)

			stationReportListener := broadcaster.Subscribe()

			for {
				currentStation := <-stationReportListener

				// Check if the passenger should board or disembark
				shouldBoard := shouldBoardPassenger(train, currentStation, passenger)
				shouldDisembark := shouldDisembarkPassenger(train, currentStation, passenger)

				var action = "ignore"
				if shouldBoard {
					action = "board"
				} else if shouldDisembark {
					action = "disembark"
				}

				passengerResponse := &PassengerResponse{
					Passenger: passenger,
					Action:    action,
				}

				if shouldBoard || shouldDisembark {
					if shouldBoard {
						// Send the passenger to the boarding channel
						train.PassengerResponseChannel <- passengerResponse
					} else {
						// Send the passenger to the disembarking channel
						train.PassengerResponseChannel <- passengerResponse
						broadcaster.CancelSubscription(stationReportListener)
						return
					}
				} else {
					// Passenger is ignoring the train report because it's not at the right station
					train.PassengerResponseChannel <- passengerResponse
				}

			}
		}(i)
	}

	go train.runTrain()

	wg.Wait()

	train.stopTrain()
}
