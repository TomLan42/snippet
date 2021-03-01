/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {

	// use a stack
	// right means pop, if does not match, report false
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if s[i] == ')' && stack[len(stack)-1] == '(' || s[i] == ']' && stack[len(stack)-1] == '[' || s[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// @lc code=end

