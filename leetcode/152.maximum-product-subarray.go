/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {

	mindp := make([]int, len(nums))
	maxdp := make([]int, len(nums))
	mindp[0] = nums[0]
	maxdp[0] = nums[0]
	result := maxdp[0]
	for i := 1; i < len(nums); i++ {
		mindp[i] = min(min(mindp[i-1]*nums[i], maxdp[i-1]*nums[i]), nums[i])
		maxdp[i] = max(max(mindp[i-1]*nums[i], maxdp[i-1]*nums[i]), nums[i])
		if result < maxdp[i] {
			result = maxdp[i]
		}
	}

	return result

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

