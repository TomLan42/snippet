/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {

	l := len(s)

	// init dp state matrix

	left, right := 0, 0
	max := 1

	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
	}

	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			left, right = i, i+1
		}
	}

	for length := 3; length < l+1; length++ {
		for i := 0; i < l-length+1; i++ {
			j := i + length - 1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true

				if length > max {
					max = length
					left = i
					right = j
				}
			}
		}
	}

	return s[left : right+1]

}

// @lc code=end

