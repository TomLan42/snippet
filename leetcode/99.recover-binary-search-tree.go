/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {

	if root == nil {
		return
	}

	// in-order tranversal
	stack := make([]*TreeNode, 0)
	visited := make([]*TreeNode, 0)
	shithole := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// visit
		visited = append(visited, root)
		if len(visited) > 1 {
			if visited[len(visited)-1].Val < visited[len(visited)-2].Val {
				shithole = append(shithole, visited[len(visited)-2], visited[len(visited)-1])
			}
		}

		// for _, node := range visited {
		// 	fmt.Printf("%d", node.Val)
		// }
		// fmt.Println("")

		// for _, node := range shithole {
		// 	fmt.Printf("%d", node.Val)
		// }
		// fmt.Println("")

		// to right
		root = root.Right
	}

	if len(shithole) >= 2 {
		temp := shithole[0].Val
		shithole[0].Val = shithole[len(shithole)-1].Val
		shithole[len(shithole)-1].Val = temp
	}

}

// @lc code=end

