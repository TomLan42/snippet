/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	mp := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-min > mp {
			mp = prices[i] - min
		}
		// update min
		if prices[i] < min {
			min = prices[i]
		}
	}

	return mp

}

// @lc code=end

