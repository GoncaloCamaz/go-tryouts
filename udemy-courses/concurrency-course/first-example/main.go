package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

// remmember, the main function itself is a go routine
func main() {
	// if we want to have this line bellow to run concurrently with main function, we can use the go keyword
	// however, if we have only go printSomething, what will happen is that the main function will exit before the printSomething function is executed
	//go printSomething("Hello World, this is the first thing to be printed")

	//printSomething("This is the second thing to be printed")
	// The lines above will print the second thing first and then the first thing because the main function will exit before the printSomething function is executed

	var wg sync.WaitGroup

	words := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota", "kappa"}

	wg.Add(len(words))

	for index, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", index, word), &wg)
	}

	wg.Wait()

	// add one wg to wait for the last printSomething to be executed
	wg.Add(1)
	printSomething("This is the last thing to be printed", &wg)
}
