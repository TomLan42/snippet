/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) int {

	result := 1
	dfs(root, &result)
	return result - 1

}

func dfs(root *TreeNode, result *int) int {
	// fidn the max path that passes this node
	// 1) root + left subtree result
	// 2) root + right subtree result
	// 3) root + left&right subtree result, but cannot return

	if root == nil {
		return 0
	}

	leftDepth := dfs(root.Left, result)
	rightDepth := dfs(root.Right, result)

	if leftDepth+rightDepth+1 > *result {
		*result = leftDepth + rightDepth + 1
	}

	return max(leftDepth, rightDepth) + 1

}

func max(x, y int) int {

	if x > y {
		return x
	}
	return y
}

// @lc code=end

