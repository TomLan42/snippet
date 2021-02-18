/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n)
	// n = 2
	dp[0] = 1

	for i := 1; i < n; i++ {
		temp := 0
		for j := 0; j < i; j++ {
			inter := max(dp[j]*dp[i-j-1], dp[j]*(i-j))
			inter = max(inter, (j+1)*dp[i-j-1])
			inter = max(inter, (j+1)*(i-j))
			if inter > temp {
				temp = inter
			}
		}
		dp[i] = temp
	}

	return dp[n-1]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

