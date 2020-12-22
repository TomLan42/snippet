/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {

	ans := []int{-1, -1}

	if len(nums) == 0 {
		return ans
	}

	// Find left most
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		ans[0] = left
	}

	// Find right most
	left = 0
	right = len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if nums[left] == target {
		ans[1] = left
	}

	return ans
}

// @lc code=end

