/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {

	flat(root)

}

func flat(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	flatten(root.Left)
	flatten(root.Right)

	//flip
	temp := root.Right
	root.Right = root.Left

	//concat
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = temp
	root.Left = nil

	return root

}

// @lc code=end

