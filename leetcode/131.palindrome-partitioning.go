/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {

	result := [][]string{}
	if len(s) == 0 {
		return result
	}
	cur := []string{}

	dfs(s, 0, cur, &result)

	return result

}

func dfs(s string, idx int, cur []string, result *[][]string) {
	start, end := idx, len(s)

	if start == end {

		temp := make([]string, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
		return
	}

	for i := start; i < end; i++ {
		if isPal(s, start, i) {
			dfs(s, i+1, append(cur, s[start:i+1]), result)
		}
	}

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

