package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "Bob"
}

func fetchUserLikes(userName string, ch chan any, wg *sync.WaitGroup) {
	fmt.Println("fetching likes for... ", userName)
	time.Sleep(time.Millisecond * 100)
	ch <- 11

	wg.Done() // decrement wg the counter
}

func fetchUserMatch(userName string, ch chan any, wg *sync.WaitGroup) {
	fmt.Println("fetching match for... ", userName)
	time.Sleep(time.Millisecond * 100)
	ch <- "Anna"

	wg.Done() // decrement wg the counter
}

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 2) // this is a buffered channel, which means, it can hold 2 values at a time before blocking the sender goroutine (fetchUserLikes and fetchUserMatch)

	wg := &sync.WaitGroup{}

	wg.Add(2) // add two goroutines to the wait group
	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)
	wg.Wait()     // block until we have two wg.Done() calls
	close(respch) // close the channel to stop the range loop

	for resp := range respch {
		fmt.Println("received:", resp)
	}

	fmt.Println("took:", time.Since(start))
}
