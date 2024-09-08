/**
* Quick Sort
* This algorithm works by selecting a pivot element and partitioning the array around the pivot, putting the pivot element in its correct position in the sorted array.
* Considering the following array: 6 3 7 5 1 2 4, we sould select the last element as the pivot.
* After the first partitioning, the array will be: 3 5 1 2 4 6 7
* The pivot element is now in its correct position. The left side of the pivot contains elements smaller than the pivot and the right side contains elements greater than the pivot.
* We then recursively sort the left and right sides of the pivot.
 */

package main

import (
	"fmt"
)

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int, low int, high int) int {
	// select the pivot element, which is located at the end of the array
	// an optimization would be to select the pivot element randomly
	pivotValue := arr[high]
	i := low
	// iterate over the array and swap elements that are smaller than the pivot
	for j := low; j < high; j++ {
		if arr[j] <= pivotValue {
			swap(arr, i, j)
			i++
		}
	}
	// swap the pivot element with the element at index i which is the correct position, since all elements to the left of i are smaller than the pivot
	swap(arr, i, high)
	return i
}

func quickSortRecursion(arr []int, low int, high int) {
	if low < high {
		// partition the array
		pivot := partition(arr, low, high)

		// sort the left side of the pivot
		quickSortRecursion(arr, low, pivot-1)
		// sort the right side of the pivot
		quickSortRecursion(arr, pivot+1, high)
	}
}

func quickSort(arr []int) {
	quickSortRecursion(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{5, 4, 3, 2, 1, 9, 8, 7, 6}
	quickSort(arr)
	fmt.Println("Sorted array:", arr)
}
