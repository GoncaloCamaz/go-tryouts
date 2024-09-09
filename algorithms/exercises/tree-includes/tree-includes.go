package main

import "fmt"

type TreeNode struct {
	nodeValue int
	left      *TreeNode
	right     *TreeNode
}

func treeIncludes(tree *TreeNode, value int) bool {
	if tree == nil {
		return false
	}
	if tree.nodeValue == value {
		return true
	}
	if value < tree.nodeValue {
		return treeIncludes(tree.left, value)
	} else {
		return treeIncludes(tree.right, value)
	}
}

func main() {
	tree := &TreeNode{nodeValue: 5, left: &TreeNode{nodeValue: 3, left: &TreeNode{nodeValue: 2}, right: &TreeNode{nodeValue: 4}}, right: &TreeNode{nodeValue: 7, left: &TreeNode{nodeValue: 6}, right: &TreeNode{nodeValue: 8}}}
	result := treeIncludes(tree, 60)
	fmt.Println("Result:", result)
}
