/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, 0, targetSum)
}

func dfs(node *TreeNode, sum, targetSum int) bool {
	sum += node.Val

	// leaf
	if node.Left == nil && node.Right == nil {
		return sum == targetSum
	}

	leftResult, rightResult := false, false
	if node.Left != nil {
		leftResult = dfs(node.Left, sum, targetSum)
	}

	if node.Right != nil {
		rightResult = dfs(node.Right, sum, targetSum)
	}

	return leftResult || rightResult
}

// @lc code=end

