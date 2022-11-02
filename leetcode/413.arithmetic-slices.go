/*
 * @lc app=leetcode id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}

	result, dp := 0, 0

	for i := 1; i < len(nums)-1; i++ {
		if (nums[i] - nums[i-1]) == (nums[i+1] - nums[i]) {
			dp++
			result += dp
		} else {
			dp = 0
		}
	}

	return result
}

// @lc code=end

