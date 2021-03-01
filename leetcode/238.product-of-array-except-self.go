/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {

	// Approach 1: product of all -> divide by num[i]

	// Approach 2: two array store product of the left and the right

	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = 1
	right[len(right)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		right[len(nums)-1-i] = right[len(nums)-i] * nums[len(nums)-i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = left[i] * right[i]
	}

	return nums

}

// @lc code=end

