/*
 * @lc app=leetcode id=1008 lang=golang
 *
 * [1008] Construct Binary Search Tree from Preorder Traversal
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
func bstFromPreorder(preorder []int) *TreeNode {

	// [root][leftroot->SubTree][righroot->SubTree]
	// leftroot < root < rightroot

	// try every plausible subtress for each node
	// backtrack?

	_, root := backtrack(preorder, 0, len(preorder)-1)
	return root

}

func backtrack(preorder []int, start, end int) (bool, *TreeNode) {

	if start > end {
		return true, nil
	}
	root := new(TreeNode)
	root.Val = preorder[start]
	if start == end {
		return true, root
	}

	// i as the end of the start of the left subtree end
	for i := start; i <= end; i++ {

		leftVal := 0
		rightVal := 0

		if start+1 > i {
			leftVal = -(1 << 31)
		} else {
			leftVal = preorder[start+1]
		}

		if i+1 > end {
			rightVal = math.MaxInt32
		} else {
			rightVal = preorder[i+1]
		}

		if leftVal < root.Val && root.Val < rightVal {
			leftOK, left := backtrack(preorder, start+1, i)
			rightOK, right := backtrack(preorder, i+1, end)

			if leftOK && rightOK {
				root.Left = left
				root.Right = right
				return true, root
			}
		}
	}
	return false, root
}

// @lc code=end

