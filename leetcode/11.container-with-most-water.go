/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {

	// two pointer

	left := 0
	right := len(height) - 1

	result := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > result {
			result = area
		}
		if height[left] == min(height[left], height[right]) {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

