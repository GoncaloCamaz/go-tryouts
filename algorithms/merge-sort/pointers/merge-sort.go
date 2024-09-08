/**
 * Merge sort is a divide-and-conquer algorithm that was invented by John von Neumann in 1945.
 * It is an efficient, stable, and comparison-based sorting algorithm.
 * Merge sort is a recursive algorithm that continually splits a list in half.
 * If the list is empty or has one item, it is sorted by definition (the base case).
 * If the list has more than one item, we split the list and recursively invoke a merge sort on each of the sublists.
 * Once the two sublists are sorted, the fundamental operation, called a merge, is performed.
 * Merging is the process of taking two smaller sorted arrays and combining them to create a single, sorted, new array.
 * This algorithm has a time complexity of O(n log n) in both best and worst cases.
 */

package main

import "fmt"

func merge(arr *[]int, left int, middle int, right int) {
	leftSize := middle - left + 1
	rightSize := right - middle

	tempLeft := make([]int, leftSize)
	tempRight := make([]int, rightSize)

	for i := 0; i < leftSize; i++ {
		tempLeft[i] = (*arr)[left+i]
	}

	for j := 0; j < rightSize; j++ {
		tempRight[j] = (*arr)[middle+1+j]
	}

	fmt.Println("Merging array: ", arr, " temp left: ", tempLeft, " temp right: ", tempRight, " left: ", left, " middle: ", middle, " right: ", right)

	for i, j, k := 0, 0, left; k <= right; k++ {
		if i < leftSize && (j >= rightSize || tempLeft[i] <= tempRight[j]) {
			(*arr)[k] = tempLeft[i]
			i++
		} else {
			(*arr)[k] = tempRight[j]
			j++
		}
	}
}

func mergeSortRecursion(arr *[]int, left int, right int) {
	if left < right {
		middle := left + (right-left)/2

		mergeSortRecursion(arr, left, middle)
		mergeSortRecursion(arr, middle+1, right)

		merge(arr, left, middle, right)
	}
}

func mergeSort(arr *[]int) {
	mergeSortRecursion(arr, 0, len((*arr))-1)
}

func main() {
	arr := []int{5, 4, 3, 2, 1, 7, 9, 8, 6}
	mergeSort(&arr)
	fmt.Println("Sorted array:", arr)
}
