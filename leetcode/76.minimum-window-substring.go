/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {

	// result realted var
	min := math.MaxInt64
	start := 0
	end := 0

	// init window
	left := 0
	right := 0
	match := 0

	// match related var
	need := make(map[byte]int)
	win := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// main logic of sliding window
	for right < len(s) {

		c := s[right]
		right++

		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		// while need to shrink
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}
			c := s[left]
			left++

			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}

		}

	}

	if min == math.MaxInt64 {
		return ""
	}

	return s[start:end]

}

// @lc code=end

