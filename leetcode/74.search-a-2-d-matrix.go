/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	left := 0
	right := m*n - 1

	for left+1 < right {
		mid := left + (right-left)/2
		val := matrix[mid/n][mid%n]

		if val > target {
			right = mid
		}
		if val < target {
			left = mid
		}
		if val == target {
			right = mid
		}
	}

	if matrix[left/n][left%n] == target || matrix[right/n][right%n] == target {
		return true
	}
	return false

}

// @lc code=end

