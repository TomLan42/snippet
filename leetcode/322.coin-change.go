/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {

	f := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		f[i] = amount + 1
	}
	f[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				f[i] = min(f[i], f[i-coins[j]]+1)
			}
		}
	}

	if f[amount] > amount {
		return -1
	}

	return f[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

