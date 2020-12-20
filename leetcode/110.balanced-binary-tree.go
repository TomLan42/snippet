/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {

	if _, isBalanced := maxDepth(root); isBalanced {
		return true
	}
	return false

}

func maxDepth(root *TreeNode) (int, bool) {

	max := 0
	isBalanced := true

	if root == nil {
		return 0, true
	}

	maxDepthLeft, isBalancedLeft := maxDepth(root.Left)
	maxDepthRight, isBalancedRight := maxDepth(root.Right)

	if maxDepthLeft-maxDepthRight <= 1 && maxDepthLeft-maxDepthRight >= -1 && isBalancedLeft && isBalancedRight {
		isBalanced = true
	} else {
		isBalanced = false
	}
	if maxDepthLeft > maxDepthRight {
		max = maxDepthLeft
	} else {
		max = maxDepthRight
	}
	return max + 1, isBalanced
}

// @lc code=end

