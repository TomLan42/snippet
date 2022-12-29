/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {

	// make sure num1 is the shorter string
	if len(num1) > len(num2) {
		return addStrings(num2, num1)
	}

	result := make([]byte, 0)

	carryDigit := 0
	for i := 0; i < len(num1); i++ {

		num1Digit := int(num1[len(num1)-1-i] - byte('0'))
		num2Digit := int(num2[len(num2)-1-i] - byte('0'))

		resultDigit := num1Digit + num2Digit + carryDigit
		carryDigit = resultDigit / 10
		resultDigit = resultDigit - carryDigit*10

		result = append(result, byte(resultDigit)+byte('0'))
	}

	for i := 0; i < len(num2)-len(num1); i++ {

		num2Digit := int(num2[len(num2)-len(num1)-1-i] - byte('0'))

		resultDigit := num2Digit + carryDigit
		carryDigit = resultDigit / 10
		resultDigit = resultDigit - carryDigit*10

		result = append(result, byte(resultDigit)+byte('0'))
	}

	if carryDigit != 0 {
		result = append(result, byte(carryDigit)+byte('0'))

	}

	realResult := make([]byte, 0)

	for i := len(result) - 1; i >= 0; i-- {
		realResult = append(realResult, result[i])
	}

	return string(realResult)

}

// @lc code=end

