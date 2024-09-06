/*
	Here we will try de depth first algorithm but with binary trees

	The idea behind depth-first search is to visit the children of a node before visiting the siblings.

	We will implement the recursive version of the depth-first search algorithm.
*/

package main

import "fmt"

type TreeNode struct {
	node  int
	left  *TreeNode
	right *TreeNode
}

func BinaryTreeDepthFirst(tree *TreeNode, visited map[int]bool, path *[]int) {
	if tree == nil {
		return
	}

	visited[tree.node] = true
	*path = append(*path, tree.node)

	BinaryTreeDepthFirst(tree.left, visited, path)
	BinaryTreeDepthFirst(tree.right, visited, path)
}

/*
The tree bellow is represented as:

		1
	   / \
	  2   3
	 / \ / \
	4  5 6  7
*/
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
	visited := map[int]bool{}
	path := []int{}

	BinaryTreeDepthFirst(tree, visited, &path)
	fmt.Println("Path followed:", path)
}
