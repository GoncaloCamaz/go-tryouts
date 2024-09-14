package main

import (
	"io"
	"os"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("Hello, Universe!", &wg)

	wg.Wait()

	if msg != "Hello, Universe!" {
		t.Errorf("Expected 'Hello, Universe!', but got: %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello, Universe!"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if output != "Hello, Universe!\n" {
		t.Errorf("Expected 'Hello, Universe!', but got: %s", output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	expected := "Hello, Universe!\nHello, Cosmos!\nHello, World!\n"

	if output != expected {
		t.Errorf("Expected %s, but got: %s", expected, output)
	}
}
