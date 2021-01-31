/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {

	left := 0
	right := len(nums) - 1
	// get rid of duplicate

	for left+1 < right {

		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid
			}
		} else if nums[mid] < nums[right] {
			if target > nums[mid] && target <= nums[right] {
				left = mid
			} else {
				right = mid
			}

		}
	}

	if nums[left] == target || nums[right] == target {
		return true
	}

	return false
}

// @lc code=end

