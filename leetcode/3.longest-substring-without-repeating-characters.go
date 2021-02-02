/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	// map related
	win := make(map[byte]int)

	// window
	left := 0
	right := 0
	dup := 0

	// result
	max := 0

	for right < len(s) {
		c := s[right]
		right++

		win[c]++
		if win[c] > 1 {
			dup++
			for dup > 0 {
				c := s[left]
				left++
				win[c]--
				if win[c] == 1 {
					dup--
				}
			}

		}

		if max < right-left {
			max = right - left
		}

	}
	return max
}

// @lc code=end

