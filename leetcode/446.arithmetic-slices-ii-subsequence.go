/*
 * @lc app=leetcode id=446 lang=golang
 *
 * [446] Arithmetic Slices II - Subsequence
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	dpMap := make([]map[int]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		dpMap[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			c := dpMap[j][diff]
			res += c
			dpMap[i][diff] += c + 1
		}
	}

	return res
}

// @lc code=end

