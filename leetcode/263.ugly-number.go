/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

// @lc code=start
func isUgly(n int) bool {

	for n != 1 && n != 0 {
		if n == 2*(n/2) {
			n = n / 2
		} else if n == 3*(n/3) {
			n = n / 3
		} else if n == 5*(n/5) {
			n = n / 5
		} else {
			break
		}
	}

	if n == 1 {
		return true
	}
	return false
}

// @lc code=end

