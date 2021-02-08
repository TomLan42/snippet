	/*
	 * @lc app=leetcode id=10 lang=golang
	 *
	 * [10] Regular Expression Matching
	 */

	// @lc code=start
	func isMatch(s string, p string) bool {
		// return recursive(s, p)
		return dpMatch(s, p)
	}

	// recursive
	func recursive(s string, p string) bool {
		// len(p) == 0
		if len(p) == 0 {
			return len(s) == 0
		}

		// len(p) >= 1
		firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
		// len(p) >= 2
		// if second is wildcard -> 0/1+
		if len(p) >= 2 && p[1] == '*' {
			return (firstMatch && recursive(s[1:], p)) || recursive(s, p[2:])
		}
		// second not a wildcard
		return firstMatch && recursive(s[1:], p[1:])
	}

	// dp
	func dpMatch(s string, p string) bool {

		m := len(s)
		n := len(p)

		// init dp array
		dp := make([][]bool, m+1)
		for i := 0; i < m+1; i++ {
			dp[i] = make([]bool, n+1)
		}
		dp[0][0] = true

		for i := 2; i <= n; i += 1 {
			if p[i-1] == '*' {
				dp[0][i] = dp[0][i-2]
			}
		}

		// start to traverse

		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if s[i-1] == p[j-1] || p[j-1] == '.' {
					dp[i][j] = dp[i-1][j-1]
				} else if p[j-1] == '*' {
					if dp[i][j-2] {
						dp[i][j] = true
					} else if s[i-1] == p[j-2] || p[j-2] == '.' {
						dp[i][j] = dp[i-1][j]
					}
				}
			}
		}

		return dp[m][n]
	}

	// @lc code=end

