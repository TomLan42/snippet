/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {

	// brute force:
	// 1 x 1 -> min(m,n) x min(m,n)
	// sliding window

	// dp
	// each entry records the value of the max len it supports as
	// the left down corner

	// init dp

	dp := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')

		}
	}

	// tranverse
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}

		}
	}

	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if dp[i][j] > result {
				result = dp[i][j]
			}

		}
	}

	return result * result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

