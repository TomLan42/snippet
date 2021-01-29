/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {

	return searchTemplate3(nums, target)

}

func searchTemplate1(nums []int, target int) []int {
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

func searchTemplate3(nums []int, target int) []int {

	result := []int{-1, -1}

	if len(nums) == 0 {
		return result
	}

	left := 0
	right := len(nums) - 1

	//find the left bound
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
	// leftover
	if nums[right] == target {
		result[0] = right
	}
	if nums[left] == target {
		result[0] = left
	}
	if nums[left] != target && nums[right] != target {
		return result
	}

	// find the right bound
	left = 0
	right = len(nums) - 1

	//find the left bound
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] < target {
			left = mid
		}
		if nums[mid] == target {
			left = mid
		}
	}
	// leftover
	if nums[left] == target {
		result[1] = left
	}
	if nums[right] == target {
		result[1] = right
	}
	if nums[left] != target && nums[right] != target {
		return result
	}

	return result

}

// @lc code=end

