/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {

	dp := make([]int, len(nums))
	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		if (i - 2) >= 0 {
			max := dp[0]
			for j := 0; j <= i-2; j++ {
				if dp[j] > max {
					max = dp[j]
				}
			}
			dp[i] = dp[i] + max
		}

		if dp[i] > result {
			result = dp[i]
		}
	}
	return result

}

// @lc code=end

