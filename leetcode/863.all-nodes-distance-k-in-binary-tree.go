/*
 * @lc app=leetcode id=863 lang=golang
 *
 * [863] All Nodes Distance K in Binary Tree
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
func distanceK(root *TreeNode, target *TreeNode, K int) []int {

	result := make([]int, 0)
	// K helps us
	findDistanceK(root, target, K, &result)
	return result

}

// find the distance from node to target
func findDistanceK(root *TreeNode, target *TreeNode, K int, result *[]int) int {

	if root == nil {
		return K + 1
	}

	// 1st type: child of target, you basically find it in the subtree
	if root == target {
		findChild(root, K, result)
		return 1
	}

	// move down a level yor K = K - 1 , if your K
	leftDistance := findDistanceK(root.Left, target, K, result)
	// if find the target within range of k
	// still got chance to extend to the right side
	// if the distance is just K, your self is the answer
	if leftDistance == K {
		findChild(root, K-leftDistance, result)
	}
	if leftDistance < K {
		findChild(root.Right, K-leftDistance-1, result)
		return leftDistance + 1
	}

	rightDistance := findDistanceK(root.Right, target, K, result)

	if rightDistance == K {
		findChild(root, K-rightDistance, result)
	}
	if rightDistance < K {
		findChild(root.Left, K-rightDistance-1, result)
		return rightDistance + 1
	}

	// nodes is found, return K+1 to let your parent do nothing
	return K + 1
}

func findChild(root *TreeNode, K int, result *[]int) {

	if root == nil {
		return
	}
	if K == 0 {
		*result = append(*result, root.Val)
	}
	if root.Left != nil {
		findChild(root.Left, K-1, result)
	}
	if root.Right != nil {
		findChild(root.Right, K-1, result)
	}
}

// @lc code=end

