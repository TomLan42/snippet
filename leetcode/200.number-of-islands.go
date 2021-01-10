/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	var count int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) > 0 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) int {
	// detect boundary
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	// paint map
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		return 1 + dfs(grid, i+1, j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
	}
	return 0
}

// @lc code=end

