/*
 * @lc app=leetcode id=329 lang=golang
 *
 * [329] Longest Increasing Path in a Matrix
 */

// @lc code=start

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func longestIncreasingPath(matrix [][]int) int {
	maxResult := 0
	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			result := dfs(matrix, cache, i, j)
			if result > maxResult {
				maxResult = result
			}
		}
	}

	return maxResult
}

func dfs(matrix [][]int, cache [][]int, i, j int) int {

	if cache[i][j] != 0 {
		return cache[i][j]
	}

	maxResult := 0
	for _, direction := range directions {
		iNext := i + direction[0]
		jNext := j + direction[1]
		if inMatrix(matrix, iNext, jNext) && matrix[iNext][jNext] > matrix[i][j] {
			result := dfs(matrix, cache, iNext, jNext)
			if result > maxResult {
				maxResult = result
			}
		}
	}

	cache[i][j] = 1 + maxResult
	return 1 + maxResult
}

func inMatrix(matrix [][]int, i, j int) bool {
	if i >= len(matrix) || i < 0 {
		return false
	}

	if j >= len(matrix[0]) || j < 0 {
		return false
	}

	return true
}

// @lc code=end

