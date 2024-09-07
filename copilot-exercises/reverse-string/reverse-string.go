package main

import "fmt"

func reverseString(s string) string {
	reversedString := make([]rune, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		reversedString[j] = rune(s[i])
	}
	return string(reversedString)
}

func main() {
	result := reverseString("Hello, World!")
	fmt.Println("Result: ", result)
}
