/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {

	//equal to find the idx of the max

	// left is from the larget half
	// right is from the smaller half
	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[right] {
			left = mid
		}
		if nums[mid] < nums[right] {
			right = mid
		}
	}

	if nums[left] < nums[right] {
		return nums[left]
	}

	return nums[right]

}

// @lc code=end

