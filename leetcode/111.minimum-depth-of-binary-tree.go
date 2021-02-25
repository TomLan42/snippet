/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	result := math.MaxInt32
	dfs(root, 1, &result)
	return result

}

func dfs(root *TreeNode, depth int, result *int) {

	if root.Left != nil {
		dfs(root.Left, depth+1, result)
	}
	if root.Right != nil {
		dfs(root.Right, depth+1, result)
	}

	if root.Right == nil && root.Left == nil {
		if depth < *result {
			*result = depth
		}
	}

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

