/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {

	// for each node at the moment, recursively
	// construct BST

	return helper(1, n)
}

func helper(start, end int) int {

	if start >= end {
		return 1
	}

	if start+1 == end {
		return 2
	}

	ans := 0
	for i := start; i <= end; i++ {
		//lefts and rights are all the roots viable of sub-trees
		lefts := helper(start, i-1)
		rights := helper(i+1, end)

		ans = ans + lefts*rights
	}

	return ans

}

// @lc code=end

