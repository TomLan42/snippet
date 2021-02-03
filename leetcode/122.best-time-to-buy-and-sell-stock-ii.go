/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {

	p := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			p += prices[i] - prices[i-1]
		}
	}
	return p

}

// @lc code=end

