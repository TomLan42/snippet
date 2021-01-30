/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left+1 < right {

		// get rid of unwanted duplicate
		for left < right && nums[left] == nums[left+1] {
			left++

		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		// binary search
		mid := left + (right-left)/2

		if nums[mid] <= nums[right] {
			right = mid

		}

		if nums[mid] > nums[right] {
			left = mid
		}
	}

	if nums[left] < nums[right] {
		return nums[left]
	} else {
		return nums[right]
	}

}

// @lc code=end

