/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {

	sum := len(nums) * (len(nums) + 1) / 2
	result := sum
	for _, num := range nums {
		result -= num
	}
	return result
}

// @lc code=end

