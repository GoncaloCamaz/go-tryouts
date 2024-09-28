package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosopher is a struct which stores information about a philosopher
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// How many times do we want the philosophers to eat before they are done
var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second // they take a lot of time to think hehe
var sleepTime = 1 * time.Second

func main() {
	// print welcome message
	fmt.Println("Dinning Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()
	// print out finished message
}

func dine() {
	// this will be the wait group for the philosophers to finish their meal
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers)) // we have 5 philosophers, so we will add them to the wait group

	// we want to wait before everyone is seated
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all five forks
	// we need this to lock the access of the philosophers to the forks
	var forks = make(map[int]*sync.Mutex)
	// initialize the forks
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire a go routine for the current philosopher
		go dinningProblem(philosophers[i], wg, forks, seated)
	}

	// wait for all philosophers to finish their meal
	wg.Wait()
}

func dinningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosohpher at the table
	fmt.Printf("%s is seated at the ttable.\n", philosopher.name)
	// decrement seated counter
	seated.Done()

	// wait for all philosophers to be seated
	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		forks[philosopher.leftFork].Lock()
		fmt.Printf("\t%s has the left fork.\n", philosopher.name)
		forks[philosopher.rightFork].Lock()
		fmt.Printf("\t%s has the right fork.\n", philosopher.name)

		fmt.Printf("\t%s is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		// release the forks
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s is done eating.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, "is done eating and satisfied.")
	fmt.Println(philosopher.name, "left the table.")

}
