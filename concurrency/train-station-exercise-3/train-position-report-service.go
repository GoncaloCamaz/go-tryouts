package main

import (
	"context"
)

type TrainPositionReportService interface {
	Subscribe() <-chan int
	CancelSubscription(<-chan int)
}

type reportService struct {
	source         <-chan int
	listeners      []chan int
	addListener    chan chan int
	removeListener chan (<-chan int)
}

func (s *reportService) Subscribe() <-chan int {
	listener := make(chan int)
	s.addListener <- listener
	return listener
}

func (s *reportService) CancelSubscription(listener <-chan int) {
	s.removeListener <- listener
}

func NewTrainReportService(ctx context.Context, source <-chan int) TrainPositionReportService {
	s := &reportService{
		source:         source,
		listeners:      make([]chan int, 0),
		addListener:    make(chan chan int),
		removeListener: make(chan (<-chan int)),
	}

	go s.serve(ctx)
	return s
}

func (s *reportService) serve(ctx context.Context) {
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
