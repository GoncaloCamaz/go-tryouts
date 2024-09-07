package main

import "fmt"

type TreeNode struct {
	node  int
	left  *TreeNode
	right *TreeNode
}

func IterativeBinaryTreeBFS(tree *TreeNode, path *[]int) {
	if tree == nil {
		return
	}

	visited := make(map[int]bool)
	queue := []*TreeNode{tree}
	for len(queue) > 0 {
		currentNode := queue[0]
		*path = append(*path, currentNode.node)
		queue = queue[1:]
		visited[currentNode.node] = true

		if currentNode.left != nil && !visited[currentNode.left.node] {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil && !visited[currentNode.right.node] {
			queue = append(queue, currentNode.right)
		}
	}
}

func main() {
	tree := &TreeNode{
		node: 1,
		left: &TreeNode{
			node: 2,
			left: &TreeNode{
				node: 4,
			},
			right: &TreeNode{
				node: 5,
			},
		},
		right: &TreeNode{
			node: 3,
			left: &TreeNode{
				node: 6,
			},
			right: &TreeNode{
				node: 7,
			},
		},
	}

	path := []int{}

	IterativeBinaryTreeBFS(tree, &path)
	fmt.Println("Path followed:", path)
}
