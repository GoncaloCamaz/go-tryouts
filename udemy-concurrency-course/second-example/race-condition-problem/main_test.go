package main

import "testing"

// we can test race conditions with go test -race .
func Test_updateMessage(t *testing.T) {
	msg = "Hello World!"

	wg.Add(2)
	go updateMessage("Hello, Universe!")
	go updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Errorf("Expected 'Goodbye, cruel world!', but got: %s", msg)
	}
}
