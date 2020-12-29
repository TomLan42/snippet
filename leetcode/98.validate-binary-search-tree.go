/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

// "validate" type of questions
// need to construct a result type to store
// bool info and other 'useful' info
type Result struct {
	IsValid bool
	Max     *TreeNode
	Min     *TreeNode
}

func helper(root *TreeNode) Result {
	result := Result{}
	if root == nil {
		result.IsValid = true
		return result
	}

	leftResult := helper(root.Left)
	rightResult := helper(root.Right)

	if !leftResult.IsValid || !rightResult.IsValid {
		result.IsValid = false
		return result
	}

	if leftResult.Max != nil && leftResult.Max.Val >= root.Val {
		result.IsValid = false
		return result
	}

	if rightResult.Min != nil && rightResult.Min.Val <= root.Val {
		result.IsValid = false
		return result
	}

	// if it is valid, consider how to report upwards
	result.IsValid = true
	result.Min = root
	if leftResult.Min != nil {
		result.Min = leftResult.Min
	}
	result.Max = root
	if rightResult.Max != nil {
		result.Max = rightResult.Max
	}
	return result

}

func isValidBST(root *TreeNode) bool {
	result := helper(root)
	return result.IsValid
}

// @lc code=end

