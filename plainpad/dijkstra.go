package plainpad

import "math"

// input will be like
// 4 3 8 9 4 5
// 6 7 8 9 0 7
// 3 0 9 8 7 4
// 6 7 8 3 9 2
// find the shortest distance from the uppper leftmost to the lower rightmost
// only the value of the next cell higher than that of the current cell is valid
// return -1 if path does not exist

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// dfs solution
func solution1(matrix [][]int) int {
	return dfs(matrix, 0, 0)
}

// return -1 if this cell cannot reach the target
func dfs(matrix [][]int, i, j int) int {
	if i == len(matrix)-1 && j == len(matrix[0])-1 {
		return 1
	}

	result := -1
	for _, direction := range directions {
		iNext := i + direction[0]
		jNext := j + direction[1]
		if matrix[iNext][jNext] > matrix[i][j] && !inMatrix(matrix, iNext, jNext) {
			continue
		} else {
			if dfs(matrix, iNext, jNext) != -1 {
				result = min(dfs(matrix, iNext, jNext), result)
			}
		}
	}

	if result != -1 {
		result++
	}

	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func inMatrix(matrix [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return false
	}
	return true
}

// Variant problem
// input will be like
// 4 3 8 9 4 5
// 6 7 8 9 0 7
// 3 0 9 8 7 4
// 6 7 8 3 9 2
// find the path from the *lowest leftmost* to the *heighest rightmost*
// that minimize

func solution2(matrix [][]int) int {
	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[0]))
	}

	return dfs2(matrix, cache, len(matrix)-1, 0)
}

func dfs2(matrix, cache [][]int, i, j int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}

	// reach the target
	if i == 0 && j == len(matrix[0])-1 {
		return matrix[0][len(matrix[0])-1]
	}

	result := math.MaxInt32
	for _, direction := range directions {
		iNext := i + direction[0]
		jNext := j + direction[1]
		if inMatrix(matrix, iNext, jNext) {
			result = min(dfs2(matrix, cache, iNext, jNext), result)
		}
	}

	cache[i][j] = result + matrix[i][j]
	return result + matrix[i][j]
}
