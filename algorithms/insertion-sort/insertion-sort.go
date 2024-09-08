/*
* Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.
* It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

	The way it works is by iterating over the list and for each element, it checks if the element is smaller than the previous element.
	Therefore, we will start the iteration from the second element and compare it with the previous one.
*/
package main

import "fmt"

func InsertionSort(arr []int) []int {
	// we start at the second element
	for i := 1; i < len(arr); i++ {
		// store the current element
		element := arr[i]
		// store the index of the previous element
		j := i - 1

		// while the index is greater than or equal to 0 and the previous element is greater than the current element

		for j >= 0 && arr[j] > element {
			// move the previous element to the next position
			arr[j+1] = arr[j]
			j--
		}
		// insert the current element in the correct position
		arr[j+1] = element
	}

	return arr
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println("Unsorted array:", arr)
	fmt.Println("Sorted array:", InsertionSort(arr))
}
