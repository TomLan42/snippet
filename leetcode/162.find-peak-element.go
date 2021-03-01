/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {

	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right-left)/2
		if mid == right || mid == left {
			break
		}
		if nums[mid] < nums[mid+1] {
			left = mid
		} else {
			right = mid
		}

	}

	if nums[left] > nums[right] {
		return left
	} else {
		return right
	}

}

// @lc code=end

