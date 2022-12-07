/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	return bsMethod(nums)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dpMethod(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	f := make([]int, len(nums))
	f[0] = 1

	ans := 0
	// f[i] indicates the longest increasing subsequece
	// to this index i

	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}

func bsMethod(nums []int) int {
	tail := make([]int, 0)

	for _, v := range nums {
		pos := searchInt(tail, v)

		if pos == len(tail) {
			tail = append(tail, v)
		} else {
			tail[pos] = v
		}

	}

	return len(tail)
}

// searchInt returns the position of the target in nums
// if not found, return the position to insert
func searchInt(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}

	if target <= nums[left] {
		return left
	}

	if target > nums[left] && target <= nums[right] {
		return right
	}

	return right + 1
}

// @lc code=end

