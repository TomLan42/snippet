/*
 * @lc app=leetcode id=395 lang=golang
 *
 * [395] Longest Substring with At Least K Repeating Characters
 */

// @lc code=start
func longestSubstring(s string, k int) int {

	max := 0

	for h := 1; h <= 26; h++ {

		left := 0
		right := 0

		freqMap := make(map[byte]int)
		uniqCount := 0
		noLessThanK := 0
		for right < len(s) {
			if uniqCount <= h {
				if freqMap[s[right]] == 0 {
					uniqCount++
				}
				freqMap[s[right]]++
				if freqMap[s[right]] == k {
					noLessThanK++
				}
				right++
			} else {
				if freqMap[s[left]] == k {
					noLessThanK--
				}
				freqMap[s[left]]--
				if freqMap[s[left]] == 0 {
					uniqCount--
				}
				left++
			}
			if uniqCount == h && noLessThanK == uniqCount && (right-left) > max {
				max = right - left
			}
		}
	}

	return max
}

// @lc code=end

