/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	left := 0
	right := len(nums) - 1

	root := constructBST(nums, left, right)
	return root
}

func constructBST(nums []int, start, end int) *TreeNode {

	if start > end {
		return nil
	}

	if start == end {

		return &TreeNode{
			Val: nums[start],
		}
	}

	mid := start + (end-start)/2
	node := new(TreeNode)

	node.Val = nums[mid]

	node.Left = constructBST(nums, start, mid-1)
	node.Right = constructBST(nums, mid+1, end)

	return node
}

// @lc code=end

