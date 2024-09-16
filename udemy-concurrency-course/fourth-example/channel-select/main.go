package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	for {
		// inside a select, if we have the same channel twice, the select will choose one of them randomly
		select {
		case s1 := <-ch1:
			fmt.Println("Case one: ", s1)
		case s2 := <-ch1:
			fmt.Println("Case two: ", s2)
		case s3 := <-ch2:
			fmt.Println("Case three: ", s3)
		case s4 := <-ch2:
			fmt.Println("Case four: ", s4)
			//default:
			// avoiding deadlock
			//	fmt.Println("No data or no go routine available")
		}
	}
}
