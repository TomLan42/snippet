/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */

// @lc code=start
func integerReplacement(n int) int {
	cache := map[]
	
	return dfs(n)
}

func dfs(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return dfs(n/2) + 1
	}

	return min(dfs(n+1), dfs(n-1)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

