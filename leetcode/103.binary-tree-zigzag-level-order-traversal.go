/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	odd := 1
	for len(queue) > 0 {
		level := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			levelEle := queue[0]
			queue = queue[1:]
			level = append(level, levelEle.Val)

			if levelEle.Left != nil {
				queue = append(queue, levelEle.Left)
			}

			if levelEle.Right != nil {
				queue = append(queue, levelEle.Right)
			}
		}
		if odd == -1 {
			level = reverse(level)
		}
		odd = odd * -1
		result = append(result, level)
	}
	return result

}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums

}

// @lc code=end

