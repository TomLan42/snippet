/*
 * @lc app=leetcode id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
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
func widthOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// use val to store how horizontal position
	dfs(root, 0)
	return bfs(root)

}

func dfs(root *TreeNode, pos int) {

	root.Val = pos
	if root.Left != nil {
		dfs(root.Left, 2*pos+1)

	}
	if root.Right != nil {
		dfs(root.Right, 2*pos+2)
	}
}

func bfs(root *TreeNode) int {

	result := 0

	// bfs of binary tree
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		l := len(queue)
		leftPos := queue[0].Val
		rightPos := queue[l-1].Val
		width := rightPos - leftPos + 1

		for i := 0; i < l; i++ {
			// record the child of those are not nil
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}

		if width > result {
			result = width
		}
	}

	return result
}

// @lc code=end

