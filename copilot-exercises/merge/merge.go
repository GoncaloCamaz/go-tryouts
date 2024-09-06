package main

import "fmt"

func merge(arr1 []int, arr2 []int) []int {
	newArraySize := len(arr1) + len(arr2)
	mergedArray := make([]int, 0, newArraySize)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArray = append(mergedArray, arr1[i])
			i++
		} else {
			mergedArray = append(mergedArray, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		mergedArray = append(mergedArray, arr1[i])
		i++
	}
	for j < len(arr2) {
		mergedArray = append(mergedArray, arr2[j])
		j++
	}
	return mergedArray
}

func main() {
	arr1 := []int{1, 3, 7, 9}
	arr2 := []int{2, 4, 6, 8}

	mergedArray := merge(arr1, arr2)
	fmt.Println(mergedArray)
}
