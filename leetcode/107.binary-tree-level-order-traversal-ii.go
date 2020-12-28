/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := make([]int, 0)
		l := len(queue)

		// dequeue 1 by 1
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
		ans = append(ans, level)
	}
	return reverse(ans)
}

// smart way to reverse a array
func reverse(result [][]int) [][]int {
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// @lc code=end

