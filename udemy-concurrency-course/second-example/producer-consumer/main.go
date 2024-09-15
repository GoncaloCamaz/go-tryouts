package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Here we will implement a producer-consumer pattern
// For that, we will try to have a pizzeria example

const numberOfPizzas = 10

var PizzasMade, PizzasFailed, Total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// important function to close the channels
func (p *Producer) close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			PizzasFailed++
		} else {
			PizzasMade++
		}

		Total++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingreditens for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// try to make pizzas
	// this cycle will run forever until we receive a quit notification
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			// selects are only useful for channels
			select {
			// we tried to make a pizza (we sent data to the data channel)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// we need to close the channels
				close(pizzaMaker.data)
				close(quitChan)

				// we want to exit the go routine
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// start the pizzeria
	go pizzeria(pizzaJob)

	// keep track of the pizzas we made
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!\n")
			}
		} else {
			color.Cyan("The Pizzeria is closed for the day")
			err := pizzaJob.close()

			if err != nil {
				color.Red("*** Error closing channel!!!", err)
			}
		}
	}

	color.Cyan("---------------------------------")
	color.Cyan("Total pizzas made: %d, but failed to make %d, with %d attempts in total.", PizzasMade, PizzasFailed, Total)

	switch {
	case PizzasFailed > 9:
		color.Red("The pizzeria is a failure!")
	case PizzasFailed >= 6:
		color.Yellow("It was not a very good day...")
	case PizzasFailed >= 4:
		color.Yellow("It was an ok day!")
	case PizzasFailed >= 2:
		color.Yellow("It was a good day!")
	default:
		color.Green("It was a great day!")
	}
}
