package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		ch <- link
		fmt.Println("Error fetching", link, ":", err)
		return
	}

	fmt.Println(link, "is up!")
	ch <- link
}

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.amazon.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	// This will block the main routine until all the goroutines are done
	// for i := 0; i < len(links)-1; i++ {
	// 	fmt.Println(<-ch)
	//}

	// This will keep the main routine running indefinitely
	for l := range ch {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, ch)
		}(l)
	}
}
