package main

import (
	"fmt"
	"sync"
)

// its important to note that the problem occurs because we are trying to acccess
// the same variable from multiple go routines
var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

func main() {
	msg = "Hello, World!"

	wg.Add(2)

	// The problem here is that we have a race condition
	// so, if we run this code with go run -race . , we will see a warning saying that we have a data race
	go updateMessage("Hello, Universe!")
	go updateMessage("Hello, Cosmos!")

	wg.Wait()

	fmt.Println(msg)
}
