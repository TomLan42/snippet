/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {

	// memoization和动规主要看子问题有没有交集

	// pure dfs: time limit exceeded
	//return DFS(0, 0, triangle)

	// dp bottom-up
	return DPBU(0, 0, triangle)
	// dp top-down

}

func DPBU(x, y int, triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	// init mem
	mem := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if mem[i] == nil {
				mem[i] = make([]int, len(triangle[i]))
			}
			mem[i][j] = triangle[i][j]
		}
	}

	// start from the second bottom
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mem[i][j] = min(mem[i+1][j], mem[i+1][j+1]) + triangle[i][j]
		}
	}

	return mem[0][0]

}

func DFS(x, y int, triangle [][]int) int {

	if x == len(triangle) {
		return 0
	}
	return min(DFS(x+1, y, triangle), DFS(x+1, y+1, triangle)) + triangle[x][y]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end



