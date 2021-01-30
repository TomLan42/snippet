/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] < target {
			left = mid
		}
		if nums[mid] == target {
			right = mid
		}
	}

	if target <= nums[left] {
		return left
	} else if target <= nums[right] {
		return right
	} else if target > nums[right] {
		return right + 1
	}

	return 0

}

// @lc code=end

