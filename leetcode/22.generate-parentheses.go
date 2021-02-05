/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {

	result := make([]string, 0)

	backtrack(n, n, "", &result)

	return result

}

// l/r are the num of available left/right parentheses
func backtrack(l, r int, str string, result *[]string) {

	if l == 0 && r == 0 {
		*result = append(*result, str)
	}

	if l > 0 {
		backtrack(l-1, r, str+"(", result)
	}
	if r > 0 && l < r {
		backtrack(l, r-1, str+")", result)
	}
}

// @lc code=end

