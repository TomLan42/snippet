package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

// recursive traversal in-order traversal
func preorderTraversal(root *Node) {

	if root == nil {
		return
	}

	fmt.Println(root.val)

	preorderTraversal(root.left)
	preorderTraversal(root.right)
}

// iterative traversal in-order traverasal

func iterativePreorderTraversal(root *Node) []int {

	result := make([]int, 0)
	stack := make([]*Node, 0)

	for root != nil || len(stack) != 0 {

		for root != nil {
			result = append(result, root.val)
			stack = append(stack, root)
			root = root.left
		}

		root = stack[len(stack)-1].right
		stack = stack[:len(stack)-1]
	}

	return result

}
