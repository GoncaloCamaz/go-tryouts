package main

import "fmt"

type TreeNode struct {
	nodeValue int
	left      *TreeNode
	right     *TreeNode
}

func treeMinValue(tree *TreeNode, minValue *int) {
	if tree == nil {
		return
	}

	if tree.nodeValue < *minValue {
		*minValue = tree.nodeValue
	}

	treeMinValue(tree.left, minValue)
	treeMinValue(tree.right, minValue)
}

func main() {
	tree := &TreeNode{nodeValue: 5, left: &TreeNode{nodeValue: 3, left: &TreeNode{nodeValue: 2}, right: &TreeNode{nodeValue: 4}}, right: &TreeNode{nodeValue: 7, left: &TreeNode{nodeValue: 6}, right: &TreeNode{nodeValue: 8}}}
	minValue := 5

	treeMinValue(tree, &minValue)

	fmt.Println("Min value:", minValue)
}
