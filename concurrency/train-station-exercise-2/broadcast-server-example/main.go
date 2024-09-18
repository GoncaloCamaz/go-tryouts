package main

import (
	"context"
	"fmt"
	"sync"
)

type BroadcastServer interface {
	Subscribe() <-chan int
	CancelSubscription(<-chan int)
}

type broadcastServer struct {
	source         <-chan int
	listeners      []chan int
	addListener    chan chan int
	removeListener chan (<-chan int)
}

func (s *broadcastServer) Subscribe() <-chan int {
	listener := make(chan int)
	s.addListener <- listener
	return listener
}

func (s *broadcastServer) CancelSubscription(listener <-chan int) {
	s.removeListener <- listener
}

func NewBroadcastServer(ctx context.Context, source <-chan int) BroadcastServer {
	s := &broadcastServer{
		source:         source,
		listeners:      make([]chan int, 0),
		addListener:    make(chan chan int),
		removeListener: make(chan (<-chan int)),
	}

	go s.serve(ctx)
	return s
}

func (s *broadcastServer) serve(ctx context.Context) {
	// close all listeners when the server is done
	defer func() {
		for _, listener := range s.listeners {
			if listener != nil {
				close(listener)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done(): // this will make the goroutine stop when the context is canceled
			return
		case newListener := <-s.addListener:
			s.listeners = append(s.listeners, newListener)
		case listenerToRemove := <-s.removeListener:
			for i, ch := range s.listeners {
				if ch == listenerToRemove {
					s.listeners = append(s.listeners[:i], s.listeners[i+1:]...)
					close(ch)
					break
				}
			}
		case val, ok := <-s.source:
			if !ok {
				return
			}
			for _, listener := range s.listeners {
				if listener != nil {
					select {
					case listener <- val:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}
}

func rangeChannel(
	ctx context.Context,
	n int,
) <-chan int {
	valueStream := make(chan int)
	go func() {
		defer close(valueStream)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done(): // this will make the goroutine stop when the context is canceled
				return
			case valueStream <- i:
			}
		}
	}()
	return valueStream
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Generates a channel sending integers
	// From 0 to 10
	range10 := rangeChannel(ctx, 10)

	broadcaster := NewBroadcastServer(ctx, range10)
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
