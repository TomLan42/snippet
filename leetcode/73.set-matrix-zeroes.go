/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {

	firstRowFlag := false
	firstColFlag := false

	// first col & row
	for i := range matrix {
		if matrix[i][0] == 0 {
			firstColFlag = true
			break
		}
	}

	for j := range matrix[0] {
		if matrix[0][j] == 0 {
			firstRowFlag = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstColFlag {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}

	if firstRowFlag {
		for j := range matrix[0] {
			matrix[0][j] = 0
		}
	}

}

// @lc code=end

