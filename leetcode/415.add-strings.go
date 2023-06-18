/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {

	// make sure num1 is the longer string
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	a := toIntSlice(num1)
	b := toIntSlice(num2)

	carry := 0
	for i := 0; i < len(a); i++ {
		if i < len(b) {
			a[i] += b[i] + carry
		} else {
			a[i] += carry
		}
		carry = a[i] / 10
		a[i] %= 10
	}

	if carry > 0 {
		a = append(a, carry)
	}

	return toString(a)
}

// to int string and reversed
func toIntSlice(num string) []int {
	result := make([]int, 0)
	for i := len(num) - 1; i >= 0; i-- {
		result = append(result, int(num[i]-'0'))
	}
	return result
}

func toString(nums []int) string {
	result := ""

	for i := range nums {
		result += string(nums[len(nums)-1-i] + '0')
	}

	return result
}

// @lc code=end

