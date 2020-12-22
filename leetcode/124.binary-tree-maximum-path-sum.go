/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

type Result struct {
	MaxSinglePath int
	MaxPath       int
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) Result {

	if root == nil {
		return Result{
			MaxSinglePath: 0,
			MaxPath:       -(1 << 31),
		}
	}

	//Divide
	leftResult := helper(root.Left)
	rightResult := helper(root.Right)

	var maxSinglePath int

	//Conquer
	if leftResult.MaxSinglePath > rightResult.MaxSinglePath {
		maxSinglePath = max(leftResult.MaxSinglePath+root.Val, 0)
	} else {
		maxSinglePath = max(rightResult.MaxSinglePath+root.Val, 0)
	}

	maxPath := max(rightResult.MaxSinglePath+leftResult.MaxSinglePath+root.Val, max(leftResult.MaxPath, rightResult.MaxPath))

	return Result{
		MaxSinglePath: maxSinglePath,
		MaxPath:       maxPath,
	}

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

