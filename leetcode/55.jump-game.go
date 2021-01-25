/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {

	if len(nums) == 1 {
		return true
	}

	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
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

