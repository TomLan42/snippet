/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {

	count := make([]int, len(citations)+1)

	for _, v := range citations {
		count[min(v, len(citations))]++
	}

	for i := len(count) - 2; i >= 0; i-- {
		count[i] += count[i+1]
		if count[i+1] >= i+1 {
			return i + 1
		}
	}

	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// @lc code=end

