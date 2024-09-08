package main

import (
	"fmt"
	"sort"
)

/**
* We are given a the root of a binary tree and a positive integer k.
* The level sum in the tree is the sum of the values of thhe nodes that are on the same level.
* Return the k-th largest level sum in the tree. If there are fewer than k levels in the tree, return -1.

* Note that two nodes are on the same level if they have the same distance from the root.
*

* Thinking process:
	So, we need to calculate the sum of the nodes at each level. For this we can use the breadth-first search algorithm since it visits the siblings before the children.
	Therefore, we can sum the values of each node at each level and store it in an array for example. Then we sort the array and get the k-th largest element.
*/

// First we need to create the TreeNode struct

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func breadthFirstSearch(tree *TreeNode, levelSums *[]int) {
	if tree == nil {
		return
	}

	visited := make(map[int]bool)
	queue := []*TreeNode{tree}

	for len(queue) > 0 {
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			visited[currentNode.value] = true

			if currentNode.left != nil && !visited[currentNode.left.value] {
				queue = append(queue, currentNode.left)
			}

			if currentNode.right != nil && !visited[currentNode.right.value] {
				queue = append(queue, currentNode.right)
			}
			sum = sum + currentNode.value
		}
		*levelSums = append(*levelSums, sum)
	}
}

func main() {
	// lets create the tree
	tree := &TreeNode{
		value: 1,
		left: &TreeNode{
			value: 2,
			left: &TreeNode{
				value: 4,
			},
			right: &TreeNode{
				value: 5,
			},
		},
		right: &TreeNode{
			value: 3,
			left: &TreeNode{
				value: 6,
			},
			right: &TreeNode{
				value: 7,
			},
		},
	}
	k := 2
	// lets create the breadth-first search algorithm and the array to store the level sums
	levelSums := []int{}
	breadthFirstSearch(tree, &levelSums)
	fmt.Println(levelSums)

	sort.Ints(levelSums)
	if levelSums[k-1] == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(levelSums[k-1])
	}

}
