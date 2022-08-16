/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
func nthUglyNumber(n int) int {

	// init dp
	dp := make([]int, n+1)
	p2, p3, p5 := 1, 1, 1
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		x2 := dp[p2] * 2
		x3 := dp[p3] * 3
		x5 := dp[p5] * 5

		x := min(x2, min(x3, x5))
		dp[i] = x
		if x == x2 {
			p2++
		}

		if x == x3 {
			p3++
		}

		if x == x5 {
			p5++
		}
	}

	return dp[n]

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

