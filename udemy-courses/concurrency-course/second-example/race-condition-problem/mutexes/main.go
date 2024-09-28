package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {
	// this will have a default value of 0
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Initial bank balance: $%d.00\n", bankBalance)
	// this will print a blank line
	fmt.Println()

	incomes := []Income{
		{Source: "Job", Amount: 5000},
		{Source: "Freelance", Amount: 200},
		{Source: "Lottery", Amount: 1000},
		{Source: "Gift", Amount: 150},
	}

	wg.Add(len(incomes))

	for index, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				// by locking and unlocking the balance, we are making sure that
				// only one go routine can access the balance at a time
				// therefore, there is no race condition here
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(index, income)
	}

	wg.Wait()

	fmt.Printf("Final bank balance: $%d.00\n", bankBalance)
}
