package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	startNow := time.Now()

	cities := []string{"London", "Lisbon", "Tokyo", "Sydney", "Paris", "Berlin", "Madrid", "Rome", "Moscow"}

	for _, city := range cities {
		data := FetchWeather(city)
		fmt.Println("This is the data", data)
	}

	fmt.Println("Done fetching weather data in:", time.Since(startNow))
}

func TestCheckWebsitesOnSteroids(t *testing.T) {
	startNow := time.Now()

	cities := []string{"London", "Lisbon", "Tokyo", "Sydney", "Paris", "Berlin", "Madrid", "Rome", "Moscow"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go FetchWeatherOnSteroids(city, ch, &wg)
	}

	// this waits for all the goroutines to finish
	go func() {
		wg.Wait()
		// we should always close the chanels
		close(ch)
	}()

	for result := range ch {
		fmt.Println("This is the result", result)
	}

	fmt.Println("Done fetching weather data in:", time.Since(startNow))
}
