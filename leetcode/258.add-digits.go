/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {

	for num > 9 {
		digitSum := 0
		for num != 0 {
			digitSum += num % 10
			num = num / 10
		}

		num = digitSum
	}

	return num
}

// @lc code=end

