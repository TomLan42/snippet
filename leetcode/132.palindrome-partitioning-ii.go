/*
 * @lc app=leetcode id=132 lang=golang
 *
 * [132] Palindrome Partitioning II
 */

// @lc code=start
func minCut(s string) int {

	// dfs all and find the result of min len -> bad idea
	// dp: min cut of the first j char -> min cut of the first i char

	if len(s) == 0 || len(s) == 1 {
		return 0
	}

	f := make([]int, len(s)+1)

	f[0] = -1
	f[1] = 0

	// i stands for "first i char"
	for i := 1; i <= len(s); i++ {
		f[i] = i - 1
		// j stands for the actual index
		for j := 0; j < i; j++ {
			if isPal(s, j, i-1) {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(s)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isPal(str string, s, e int) bool {
	for s < e {
		if str[s] != str[e] {
			return false
		}
		s++
		e--
	}
	return true
}

// @lc code=end

