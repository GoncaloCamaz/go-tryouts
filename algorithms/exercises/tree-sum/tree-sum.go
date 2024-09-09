package main

type TreeNode struct {
	nodeValue int
	left      *TreeNode
	right     *TreeNode
}

func treeSum(tree *TreeNode, sum int) int {
	if tree == nil {
		return sum
	} else {
		return tree.nodeValue + treeSum(tree.left, sum) + treeSum(tree.right, sum)
	}
}

func main() {
	tree := &TreeNode{nodeValue: 5, left: &TreeNode{nodeValue: 3, left: &TreeNode{nodeValue: 2}, right: &TreeNode{nodeValue: 4}}, right: &TreeNode{nodeValue: 7, left: &TreeNode{nodeValue: 6}, right: &TreeNode{nodeValue: 8}}}
	result := treeSum(tree, 0)
	println("Result:", result)
}
