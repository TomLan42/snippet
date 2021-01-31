/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {

	// for each node at the moment, recursively
	// construct BST

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {

	if start > end {
		return []*TreeNode{nil}
	}

	ans := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		//lefts and rights are all the roots viable of sub-trees
		lefts := helper(start, i-1)
		rights := helper(i+1, end)
		for k := 0; k < len(lefts); k++ {
			for j := 0; j < len(rights); j++ {
				root := new(TreeNode)
				root.Val = i
				root.Left = lefts[k]
				root.Right = rights[j]
				ans = append(ans, root)
			}
		}
	}
	return ans

}

// @lc code=end

