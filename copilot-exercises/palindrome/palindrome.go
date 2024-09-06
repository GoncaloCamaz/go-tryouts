package main

import (
	"fmt"
	"regexp"
	"strings"
)

func cleanString(s string) string {
	cleanedStr := regexp.MustCompile("[^a-zA-Z]+").ReplaceAllString(s, "")
	return strings.ToLower(cleanedStr)
}

func isPalindrome(s string) bool {
	cleanedString := cleanString(s)
	for i, j := 0, len(cleanedString)-1; i < j; i, j = i+1, j-1 {
		if rune(cleanedString[j]) != rune(cleanedString[i]) {
			return false
		}
	}
	return true
}

func isPalindromeWithPointers(s *string) bool {
	cleanedString := cleanString(*s)
	for i, j := 0, len(cleanedString)-1; i < j; i, j = i+1, j-1 {
		if rune(cleanedString[j]) != rune(cleanedString[i]) {
			return false
		}
	}
	return true
}

func main() {
	// Write a function that determines if a string is a palindrome
	// A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward
	// For example, "madam" is a palindrome
	// The function should return true if the string is a palindrome, otherwise it should return false
	// The function should ignore spaces and capitalization

	testString := "Mr. Owl ate my metal worm"
	fmt.Println("The test string is a palindrome: ", isPalindrome(testString))

	// Call the isPalindromeWithPointers function by passing a pointer to the testString variable
	fmt.Println("The test string is a palindrome (with pointers): ", isPalindromeWithPointers(&testString))
}
