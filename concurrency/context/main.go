package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10

	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result", val)
	fmt.Println("took:", time.Since(start))
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyDependencies()
		respch <- Response{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("timeout fetching user: %d", userID)
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

// This will be a function that fetches some third party dependency
// and it might take a while to complete
func fetchThirdPartyDependencies() (int, error) {
	time.Sleep(time.Millisecond * 300)

	return 666, nil
}
