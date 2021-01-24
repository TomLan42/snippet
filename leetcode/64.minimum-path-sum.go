/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// init
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}

	// state evolve
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

