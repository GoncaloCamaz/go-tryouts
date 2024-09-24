package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

/** There are several ways to traverse a tree. Bellow we will implement the following:
 * TraverseTreeInOrder traverses the tree in order
 * TraverseTreePostOrder traverses the tree post order
 * TraverseTreePreOrder traverses the tree pre order
 * TraverseTreeLevelOrder traverses the tree level order
 */

// In order traversal is a depth-first traversal that starts at the root node and traverses the left subtree first, then the root node, and finally the right subtree.
func TraverseTreeInOrder(tree *TreeNode, inOrderPath *[]int) {
	if tree == nil {
		return
	}

	TraverseTreeInOrder(tree.Left, inOrderPath)
	*inOrderPath = append(*inOrderPath, tree.Value)
	TraverseTreeInOrder(tree.Right, inOrderPath)
}

// Post order traversal is a depth-first traversal that starts at the root node and traverses the left subtree first, then the right subtree, and finally the root node.
func TraverseTreePostOrder(tree *TreeNode, postOrderPath *[]int) {
	if tree == nil {
		return
	}

	TraverseTreePostOrder(tree.Left, postOrderPath)
	TraverseTreePostOrder(tree.Right, postOrderPath)
	*postOrderPath = append(*postOrderPath, tree.Value)
}

// Pre order traversal is a depth-first traversal that starts at the root node and traverses the root node first, then the left subtree, and finally the right subtree.
func TraverseTreePreOrder(tree *TreeNode, preOrderPath *[]int) {
	if tree == nil {
		return
	}

	*preOrderPath = append(*preOrderPath, tree.Value)
	TraverseTreePreOrder(tree.Left, preOrderPath)
	TraverseTreePreOrder(tree.Right, preOrderPath)
}

// Level order traversal is a breadth-first traversal that starts at the root node and traverses the tree level by level.
func TraverseTreeLevelOrder(tree *TreeNode, levelOrderPath *[]int) {
	if tree == nil {
		return
	}

	queue := []*TreeNode{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		*levelOrderPath = append(*levelOrderPath, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	tree := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left: &TreeNode{
				Value: 4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Value: 3,
			Left: &TreeNode{
				Value: 6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Value: 7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	inOrderPath := []int{}
	TraverseTreeInOrder(tree, &inOrderPath)
	fmt.Println("in order path", inOrderPath)

	postOrderPath := []int{}
	TraverseTreePostOrder(tree, &postOrderPath)
	fmt.Println("post order path", postOrderPath)

	preOrderPath := []int{}
	TraverseTreePreOrder(tree, &preOrderPath)
	fmt.Println("pre order path", preOrderPath)

	levelOrderPath := []int{}
	TraverseTreeLevelOrder(tree, &levelOrderPath)
	fmt.Println("level order path", levelOrderPath)
}
