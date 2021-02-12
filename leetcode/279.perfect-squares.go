/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {

	dp := make([]int, n+1)

	anchor := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		anchor = append(anchor, i*i)
	}

	// iterate dp array
	for i := 1; i <= n; i++ {
		ans := i
		for j := len(anchor) - 1; j >= 0; j-- {
			if i-anchor[j] >= 0 {
				if dp[i-anchor[j]] < ans {
					dp[i] = dp[i-anchor[j]] + 1
					ans = dp[i-anchor[j]]
				}
			}
		}
	}

	return dp[n]
}

// @lc code=end

