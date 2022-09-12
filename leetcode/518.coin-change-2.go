/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 */

// @lc code=start
func change(amount int, coins []int) int {

	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// @lc code=end

