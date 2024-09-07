/*
	Here, we will implement the depth-first search algorithm in Go.

	Depth-first search is an algorithm used to traverse and search tree or graph data structures.
	It starts at the root node and explores as far as possible along each branch before backtracking.

	There are two ways to implement depth-first search: recursively and iteratively.

	Let's say we have the following graph:

		0
	   / \
	  1   2
	 / \   \
	3   4   5

	We can represent this graph as an adjacency list:

	0: [1, 2]
	1: [0, 3, 4]
	2: [0, 5]
	3: [1]
	4: [1]
	5: [2]

	We will implement the recursive version of the depth-first search algorithm.

	Here is the implementation:
*/

package main

import (
	"fmt"
)

func GraphRecursiveDepthFirst(graph map[int][]int, vertex int, visited map[int]bool, path *[]int) {
	// mark the current vertex as visited
	visited[vertex] = true
	neighbors := graph[vertex]

	// range returns the index and the value of the slice, i, neighbor := range neighbors
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			*path = append(*path, neighbor)
			GraphRecursiveDepthFirst(graph, neighbor, visited, path)
		}
	}
}

func GraphRecursiveDepthFirstPostOrder(graph map[int][]int, vertex int, visited map[int]bool, path *[]int) {
	visited[vertex] = true
	neighbors := graph[vertex]

	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			GraphRecursiveDepthFirstPostOrder(graph, neighbor, visited, path)
		}
	}

	*path = append(*path, vertex)
}

func GraphIterativeDepthFirst(graph map[int][]int, startVertex int, path *[]int) {
	visited := make(map[int]bool)

	stack := []int{startVertex}
	visited[startVertex] = true

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		neighbors := graph[vertex]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack = append(stack, neighbor)
				*path = append(*path, neighbor)
			}
		}
	}
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 5},
		3: {1},
		4: {1},
		5: {2},
	}
	visited := make(map[int]bool)

	path := make([]int, 0, len(graph))

	GraphRecursiveDepthFirst(graph, 0, visited, &path)
	fmt.Println("Path:", path)
}
