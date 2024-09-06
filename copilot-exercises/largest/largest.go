package main

import "fmt"

func getLargestElement(arr []int) int {
	largest := arr[0]

	for _, value := range arr {
		if value > largest {
			largest = value
		}
	}

	return largest
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := getLargestElement(array)
	fmt.Println("Largest element: ", result)
}
