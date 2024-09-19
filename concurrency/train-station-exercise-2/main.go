package main

import (
	"context"
	"fmt"
	"sync"
)

var totalStops = 5
var startingStop = 0

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	train := NewTrain(totalStops, true)

	broadcaster := NewTrainReportService(ctx, train.CurrentStopChannel)
	listener1 := broadcaster.Subscribe()
	listener2 := broadcaster.Subscribe()
	listener3 := broadcaster.Subscribe()

	var wg sync.WaitGroup
	// we will have 3 listeners, so we should wait for 3 goroutines
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := range listener1 {
			fmt.Printf("Listener 1: %v/10 \n", i+1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range listener2 {
			fmt.Printf("Listener 2: %v/10 \n", i+1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range listener3 {
			fmt.Printf("Listener 3: %v/10 \n", i+1)
		}
	}()
	wg.Wait()
}
