package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func printInOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	printInOrder(tree.Left)
	fmt.Printf("%d ", tree.Value)
	printInOrder(tree.Right)
}

func mirror(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}

	left := mirror(tree.Left)
	right := mirror(tree.Right)

	tree.Left = right
	tree.Right = left

	return tree
}

func main() {
	tree := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 4,
			},
			Right: &TreeNode{
				Value: 5,
			},
		},
		Right: &TreeNode{
			Value: 3,
			Left: &TreeNode{
				Value: 6,
			},
			Right: &TreeNode{
				Value: 7,
			},
		},
	}

	fmt.Printf("Original tree: %+v\n", tree)
	printInOrder(tree)
	fmt.Println()

	mirrorTree := mirror(tree)
	fmt.Printf("Mirrored tree: %+v\n", mirrorTree)
	printInOrder(mirrorTree)
	fmt.Println()
}
