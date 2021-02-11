/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// inorder
	// [L ][Root][R ], [L][Root], [Root][R]
	// postorder
	// [L ][R ][Root], [L][Root], [R][Root]
	// find root
	// root.Left = xxx
	// root.Right = xxx
	inPos := make(map[int]int)

	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return dfs(inorder, postorder, 0, len(postorder)-1, 0, inPos)
}

func dfs(inorder, postorder []int, postStart, postEnd, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{
		Val: postorder[postEnd],
	}
	inRoot := inPos[postorder[postEnd]]
	leftLen := inRoot - inStart

	root.Left = dfs(inorder, postorder, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = dfs(inorder, postorder, postStart+leftLen, postEnd-1, inRoot+1, inPos)
	return root
}

// @lc code=end

