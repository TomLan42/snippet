/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {

	// store intermediate result
	f := make([]int, len(s)+1)
	f[0] = 1
	if s[0] == '0' {
		f[1] = 0
	} else {
		f[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		if s[i-1] >= '1' && s[i-1] <= '9' {
			f[i] += f[i-1]

		}
		if s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			f[i] += f[i-2]
		}

	}
	return f[len(s)]
}

// @lc code=end

