/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {

	// return dp(height)
	return twoPointer(height)

}

// dp
func dp(height []int) int {

	if len(height) == 0 {
		return 0
	}

	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))

	maxLeft[0] = height[0]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	result := 0
	for i := 1; i < len(height)-1; i++ {
		result += min(maxLeft[i], maxRight[i]) - height[i]
	}
	return result

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// two-pointer

func twoPointer(height []int) int {

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	for left <= right {
		// find minimum of both sides
		if height[left] <= height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}

			left++

		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}

// @lc code=end

