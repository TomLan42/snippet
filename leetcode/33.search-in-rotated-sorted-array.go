/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// left half
		if nums[mid] > nums[left] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid
			}
		} else if nums[mid] < nums[right] {
			if target >= nums[mid] && target <= nums[right] {
				left = mid
			} else {
				right = mid
			}
		}
	}

	if nums[left] == target {
		return left
	} else if nums[right] == target {
		return right
	}

	return -1
}

// @lc code=end

