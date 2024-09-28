package main

import "testing"

/*
%d: Decimal integer
%b: Binary integer
%o: Octal integer
%x: Hexadecimal integer (lowercase)
%X: Hexadecimal integer (uppercase)
%f: Floating-point number
%e: Scientific notation (lowercase)
%E: Scientific notation (uppercase)
%s: String
%q: Double-quoted string
%v: Default format for the value
%T: Type of the value
%p: Pointer address
/*

* TestHello() function is used to test the Hello() function
* t *testing.T is used to print the error message
* Every test must be in a file with a name like xxx_test.go
* The test function must start with the word Test
* The test function takes one argument only t *testing.T
*/
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Camaz")
		want := "Hello, Camaz"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

/*
In Go, the testing.TB interface is used as a parameter in test functions to provide a way to interact with the testing framework.
The TB interface is actually an interface that is implemented by both *testing.T and *testing.B types.

The *testing.T type represents a test case, while the *testing.B type represents a benchmark case.
Both types implement the TB interface, which means that they can be used interchangeably as the t parameter in test functions.

By using the testing.TB interface instead of the specific *testing.T type, it allows the test functions to be used as both regular tests and benchmarks without any modifications.
This provides flexibility and allows the code to be reused in different testing scenarios.

So, when you see t testing.TB as a parameter in a test function, it means that the function can be used as both a regular test and a benchmark, and it can accept either a *testing.T or a *testing.B instance.
*/
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // is needed to say that this is an helper/util function
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Another important notes, in GO, public functions start with a capital letter and private functions start with a lowercase letter.
