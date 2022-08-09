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

// iterative traversal in-order

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

// iterative traversal in-order

func iterativeInOrderTraversal(root *Node) []int {

	result := make([]int, 0)
	stack := make([]*Node, 0)

	for root != nil || len(result) > 0 {

		// left
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		// right
		node := stack[len(stack)-1]
		result = append(result, node.val)
		stack = stack[:len(stack)-1]
		root = node.right
	}

	return result
}

// iterative traversal post-order

func iterativePostOrderTraversal(root *Node) []int {
	result := make([]int, 0)
	stack := make([]*Node, 0)

	var lastVisit *Node

	for root != nil || len(result) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		node := stack[len(stack)-1]

		// stack = stack[:len(stack)-1]
		if node.right == nil || lastVisit == node.right {
			stack = stack[:len(stack)-1]
			result = append(result, node.val)
			lastVisit = node
		} else {
			root = node.right
		}

	}

}
