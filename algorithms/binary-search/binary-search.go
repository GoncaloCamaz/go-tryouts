/*
	For binary search, we must have a sorted array
	The idea behing binary search is that, the array being sorted,
	if the number we want is greater than the middle element, we can discard the left half of the array or vice versa if the number is smaller than the middle element.

	This can be implemented recursevely or iteratively.

	[2, 4, 6, 8, 10, 12, 14, 16, 18, 20]

	Let's say we want to find the number 12 in the array above.

	1. We start by finding the middle element of the array, which is 10.
	2. Since 12 is greater than 10, we can discard the left half of the array.
	3. We now have [12, 14, 16, 18, 20]
	4. We repeat the process until we find the number we want or we have no more elements to search.

	This algorithm has a time complexity of O(log n) since we are dividing the array in half at each iteration.
*/

package main

func IterativeBinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	// While both ends of the array do not meet
	for low <= high {
		// calculate middle element
		mid := low + (high-low)/2

		// if the middle element is the target, return the index
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			// if the middle element is less than the target, discard the left half of the array
			low = mid + 1
		} else {
			// if the middle element is greater than the target, discard the right half of the array
			high = mid - 1
		}
	}

	return -1
}

func RecursiveBinarySearch(arr []int, target, low, high int) int {
	// if low is greater than high, the target is not in the array
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return RecursiveBinarySearch(arr, target, mid+1, high)
	} else {
		return RecursiveBinarySearch(arr, target, low, mid-1)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 18

	index := RecursiveBinarySearch(arr, target, 0, len(arr)-1)
	if index != -1 {
		println("Target found at index: ", index)
		println(1 + (6-1)/2)
	} else {
		println("Target not found")
	}
}
