package main

import (
	"math/rand"
	"time"
)

type Ticket struct {
	Price       float64
	Origin      int
	Destination int
}

type Passenger struct {
	Name   string
	Ticket Ticket
}

type PassengerResponse struct {
	Action    string // literals: board, disembark, ignore
	Passenger *Passenger
}

func NewPassenger(name string, ticket Ticket) *Passenger {
	return &Passenger{
		Name:   name,
		Ticket: ticket,
	}
}

func generatePassengerTicket() Ticket {
	rand.Seed(time.Now().UnixNano())

	possibleStations := make([]int, 0)
	for i := 0; i < totalStops; i++ {
		possibleStations = append(possibleStations, i)
	}

	origin := possibleStations[rand.Intn(len(possibleStations))]
	destination := possibleStations[rand.Intn(len(possibleStations))]
	for origin == destination {
		destination = possibleStations[rand.Intn(len(possibleStations))]
	}

	return Ticket{
		Price:       10,
		Origin:      origin,
		Destination: destination,
	}
}
