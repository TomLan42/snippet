/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {

	// window
	left := 0
	right := 0
	match := 0

	// map related
	need := make(map[byte]int)
	win := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	for right < len(s2) {

		c := s2[right]
		right++

		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		for right-left >= len(s1) {
			if match == len(need) {
				return true
			}

			c := s2[left]
			left++

			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}

		}

	}
	return false

}

// @lc code=end

