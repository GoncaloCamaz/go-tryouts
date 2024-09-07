/*
	Here we will implement the breadth-first search algorithm for graphs.

	The idea behind breadth-first search is to visit the siblings of a node before visiting the children.

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

	So, bfs here will start by visiting 0 and then its neighbours (1 and 2).
	Then it will visit 1 and 2 neighbours (3, 4 and 5).
*/

package main

import "fmt"

func GraphIterativeBreadthFirst(graph map[int][]int, startVertex int, path *[]int) {
	visited := make(map[int]bool)
	queue := []int{startVertex}

	for len(queue) > 0 {
		currentVertex := queue[0]
		visited[currentVertex] = true
		queue = queue[1:]
		*path = append(*path, currentVertex)

		for _, neighbour := range graph[currentVertex] {
			if !visited[neighbour] {
				queue = append(queue, neighbour)
				visited[neighbour] = true
			}
		}
	}
}

func GraphRecursiveBreadthFirst(graph map[int][]int, vertex int, visited map[int]bool, queue []int, path *[]int) {
	visited[vertex] = true
	*path = append(*path, vertex)
	queue = append(queue, graph[vertex]...)

	for _, neighbour := range queue {
		if !visited[neighbour] {
			queue = queue[1:]
			GraphRecursiveBreadthFirst(graph, neighbour, visited, queue, path)
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
	path := make([]int, 0, len(graph))

	GraphIterativeBreadthFirst(graph, 0, &path)
	fmt.Println("Iterative path:", path)

	path = []int{}
	visited := make(map[int]bool)
	queue := []int{}
	GraphRecursiveBreadthFirst(graph, 0, visited, queue, &path)
	fmt.Println("Recursive path:", path)
}
