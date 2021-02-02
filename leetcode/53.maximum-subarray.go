/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {

	max := math.MinInt64
	// init f
	f := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		f[i] = nums[i-1]
		if f[i]+f[i-1] > f[i] {
			f[i] += f[i-1]
		}
		if f[i] > max {
			max = f[i]
		}
	}

	return max

}

// @lc code=end

