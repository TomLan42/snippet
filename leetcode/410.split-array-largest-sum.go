/*
 * @lc app=leetcode id=410 lang=golang
 *
 * [410] Split Array Largest Sum
 */

// @lc code=start
func splitArray(nums []int, k int) int {
	minBound := math.MinInt32
	maxBound := 0

	for _, num := range nums {
		maxBound += num
		minBound = max(minBound, num)
	}

	// binary search
	for minBound < maxBound {
		mid := minBound + (maxBound-minBound)/2
		if splittable(nums, mid, k) {
			maxBound = mid
		} else {
			minBound = mid + 1
		}
	}

	return maxBound
}

func splittable(nums []int, subArrayUpperBound, k int) bool {

	subArraySum := 0
	subArrayCount := 1

	for _, num := range nums {
		subArraySum += num
		if subArraySum > subArrayUpperBound {
			subArrayCount++
			subArraySum = num
			if subArrayCount > k {
				return false
			}
		}
	}

	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

