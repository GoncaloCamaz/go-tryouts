package main

import (
	"fmt"
	"strings"
)

// using ping <-chan string says that ping is a read only channel
// using pong chan<- string says that pong is a write only channel
// this is a good practice to make sure that the channels are used correctly
func shout(ping <-chan string, pong chan<- string) {
	// ping will be the channel that receives data and pong will be the channel that sends data
	for {
		s := <-ping

		if s == "mahnamahna" {
			pong <- "Do doo be-do-do"
		} else {
			pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
		}
	}
}

func main() {
	// create two channels, they only accept string
	ping := make(chan string)
	pong := make(chan string)

	// start the shout goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done! Closing channels")
	close(ping)
	close(pong)
}
