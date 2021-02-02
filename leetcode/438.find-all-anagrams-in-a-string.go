/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {

	// window
	left := 0
	right := 0
	match := 0

	// map ralated
	need := make(map[byte]int)
	win := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	// result
	result := make([]int, 0)

	for right < len(s) {
		c := s[right]
		right++

		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		for right-left >= len(p) {
			// if there's a match
			if right-left == len(p) && match == len(need) {
				result = append(result, left)
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
	return result

}

// @lc code=end

