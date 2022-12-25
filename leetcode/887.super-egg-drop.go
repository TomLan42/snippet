/*
 * @lc app=leetcode id=887 lang=golang
 *
 * [887] Super Egg Drop
 */

// @lc code=start
func superEggDrop(k int, n int) int {
	// cache := make([][]int, k+1)
	// for i := range cache {
	// 	cache[i] = make([]int, n+1)
	// }
	// return helper(k, n, cache)

	return dp(k, n)
}

func helper(k int, n int, cache [][]int) int {
	if cache[k][n] != 0 {
		return cache[k][n]
	}

	if k == 1 || n == 1 || n == 0 {

		return n
	}

	minResult := math.MaxInt
	for i := 1; i <= n; i++ {
		// 1. if break. test the lower half floor with one less egg.
		// 2. if not break. test the upper half floor with k egg.
		worstOutCome := 1 + max(helper(k-1, i-1, cache), helper(k, n-i, cache))
		if worstOutCome < minResult {
			minResult = worstOutCome
		}
	}

	cache[k][n] = 1 + minResult
	return 1 + minResult
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func dp(k, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][1] = 1
	}

	for i := range dp[1] {
		dp[1][i] = i
	}

	for i := 2; i <= k; i++ {
		for j := 2; j <= n; j++ {
			minResult := math.MaxInt
			for h := 1; h <= j; h++ {
				minResult = min(minResult, max(dp[i-1][h-1], dp[i][j-h]))
			}
			dp[i][j] = 1 + minResult
		}
	}

	return dp[k][n]
}

// @lc code=end

