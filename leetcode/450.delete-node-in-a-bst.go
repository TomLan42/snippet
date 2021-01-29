/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
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
func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return root
	}
	// search for the key
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		if root.Right != nil && root.Left != nil {
			cur := root.Right

			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}

// @lc code=end

