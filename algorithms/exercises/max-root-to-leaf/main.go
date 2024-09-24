package main

/**
 * Given a binary tree, find the maximum path sum. The path may start and end at any node in the tree.
 * For example, given the below binary tree,
 *        1
 *       / \
 *      2   3
 * Return 6.
 */

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func MaxPathSum(tree *TreeNode, maxPathSum *int) int {
	if tree == nil {
		return 0
	}

	left := MaxPathSum(tree.Left, maxPathSum)
	right := MaxPathSum(tree.Right, maxPathSum)

	*maxPathSum = max(*maxPathSum, left+right+tree.Value)

	return max(left, right) + tree.Value
}

func main() {
	tree := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Value: 3,
			Left:  nil,
			Right: nil,
		},
	}

	maxPathSum := 0
	MaxPathSum(tree, &maxPathSum)
	println("Max path sum:", maxPathSum)
}
