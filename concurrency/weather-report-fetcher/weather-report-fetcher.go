package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const apiKey = "bd5e378503939ddaee76f12ad7a97608"

func FetchWeather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather data: %s: %s\n", city, err)
		return data
	}

	// defer will always be executed before the function returns
	defer resp.Body.Close()

	// get error from response, decode it, and if there is an error, print it
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data: %s: %s\n", city, err)
		return data
	}

	return data
}

func FetchWeatherOnSteroids(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather data: %s: %s\n", city, err)
		return data
	}

	// defer will always be executed before the function returns
	defer resp.Body.Close()

	// get error from response, decode it, and if there is an error, print it
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data: %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf(("This is the data for %s: %v"), city, data)

	return data
}
