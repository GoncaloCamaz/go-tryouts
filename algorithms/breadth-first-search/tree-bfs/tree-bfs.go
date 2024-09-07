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

func RecursiveBinaryTreeBFS(tree *TreeNode, visited map[int]bool, queue []*TreeNode, path *[]int) {
	// stop recursion if tree is nil
	// this may happen when we reach a leaf node
	if tree == nil {
		return
	}

	visited[tree.node] = true
	*path = append(*path, tree.node)
	// append the left and right nodes to the queue
	queue = append(queue, tree.left, tree.right)

	// iterate over the queue
	for _, treeNode := range queue {
		if treeNode != nil && !visited[treeNode.node] {
			// remove the first element from the queue
			queue = queue[1:]
			// recursively call the function with the current node
			RecursiveBinaryTreeBFS(treeNode, visited, queue, path)
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
	path = []int{}
	RecursiveBinaryTreeBFS(tree, map[int]bool{}, []*TreeNode{tree}, &path)
	fmt.Println("Path followed:", path)
}
