package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	// pls dont forget to defer the wg.Done() function to avoid deadlock
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {
	// challange: modify this code so that the calls to updateMessage()
	// as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage()
	// printMessage(), and main().

	msg = "Hello, World!"

	wg.Add(1)
	go updateMessage("Hello, Universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, Cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, World!", &wg)
	wg.Wait()
	printMessage()
}
