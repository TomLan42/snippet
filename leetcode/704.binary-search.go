/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid
		}
		if target > nums[mid] {
			left = mid
		}
		if target == nums[mid] {
			return mid
		}
	}
	// two left over
	if target == nums[left] {
		return left
	}
	if target == nums[right] {
		return right
	}
	return -1
}

// @lc code=end

