package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got data:", i)

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// in the code bellow, we are creating an unbuffered channel. This is blockimg, which means that the sender will block until the receiver is ready to receive the data,
	// in this case, the listen to chan function
	ch := make(chan int)
	// if we want to create a buffered channel, we can do it like this:
	//ch := make(chan int, 10)

	// if we use a buffered channel, with 10, for example, the sender will block until the receiver is ready to receive the data, but the receiver will not block until the buffer is full
	// so we will have 10 prints of data being sent at a time

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending data:", i)
		ch <- i
		fmt.Println("sent", i, "to channel")
	}

	fmt.Println("All done!")
	close(ch)
}
