/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		// level is the answer for every level
		level := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			// dequeue
			levelEle := queue[0]
			queue = queue[1:]
			level = append(level, levelEle.Val)

			// enqueue
			if levelEle.Left != nil {
				queue = append(queue, levelEle.Left)
			}
			if levelEle.Right != nil {
				queue = append(queue, levelEle.Right)
			}
		}

		ans = append(ans, level)
	}
	return ans
}

// @lc code=end

