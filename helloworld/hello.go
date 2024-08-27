package main

import "fmt"

// 1 - first run
/*
* main functions is the entry point of the program
* fmt package is used to print the output
* fmt.Println() is used to print the output
* "Hello, World!" is the output
 */
// func main() {
// 	fmt.Println("Hello, World!")
// }

// 2 - calling a function
// func Hello() string {
// 	return "Hello, World!"
// }

// func main() {
// 	fmt.Println(Hello())
// }

// 3 - refactoring to accept strings as inputs
func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("World!"))
}
