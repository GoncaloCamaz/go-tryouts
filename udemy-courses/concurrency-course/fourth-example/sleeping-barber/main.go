package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

/**
Problem statement

This is a simulation of the sleeping barber problem. The problem is as follows:

Here we have a finite number of barbers, a finite number of seats in the waiting room, a fixed length of time the barbershop is open, and clients arriving at regular intervals.
When a barber has nothing to do, he or she checks the waiting rom for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to sleep until a new client arrives. So the rules are as follows:

- if there are no customers, the barber falls asleep in the chair
- a customer must wake the barber if he is asleep
- if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and sits in an empty chair if its available
- when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers and falls asleep if there are none
- shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is empty
- after the shop is closed and there are no clients left in the waiting area, the barber goes home

*/

// variables
var seatingCapacity = 10
var clientArrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())
	// print a welcome message
	color.Yellow("The Sleeping Barber problem")
	color.Yellow("----------------------------")

	// create channels if we need any
	// we must have a chan to send clients to...
	// we are going to use a buffered channel to simulate the waiting room with a fixed number of seats, in this case, seatingCapacity
	clientChannel := make(chan string, seatingCapacity)
	doneChannel := make(chan bool)

	// create data struct to represent the barbershop
	shop := BarberShop{
		ShopCapacity:       seatingCapacity,
		HairCutDuration:    cutDuration,
		NumberOfBarbers:    0,
		ClientsChannel:     clientChannel,
		BarbersDoneChannel: doneChannel,
		IsOpen:             true,
	}

	color.Green("The barbershop is now open!")

	// add barbers
	// we are going to use a go routine to simulate the barbers
	shop.addBarber("Frank")
	shop.addBarber("Gerard")
	shop.addBarber("Bruno")
	shop.addBarber("Milton")
	shop.addBarber("Susan")

	// start the barbershop as a go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * clientArrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed

	time.Sleep(5 * time.Second)
}
