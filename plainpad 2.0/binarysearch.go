package main

// binary search, post-processing last 2 candidates
func binarySearch1(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left+1 < right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}

	}

	if nums[left] == target {
		return left
	}

	if nums[right] == right {
		return right
	}

	return -1

}

//binary search, no post-processing
func binarySearch2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1

}
