/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
func maxProfit(prices []int) int {

	// bi-directional dp
	leftProfit := make([]int, len(prices))
	rightProfit := make([]int, len(prices)+1)

	mp := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > mp {
			mp = prices[i] - min
			leftProfit[i] = mp
		} else {
			leftProfit[i] = leftProfit[i-1]
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	mp = 0
	max := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if max-prices[i] > mp {
			mp = max - prices[i]
			rightProfit[i] = mp
		} else {
			rightProfit[i] = rightProfit[i+1]
		}
		if prices[i] > max {
			max = prices[i]
		}
	}

	result := 0

	for i := 0; i < len(prices); i++ {
		if leftProfit[i]+rightProfit[i+1] > result {
			result = leftProfit[i] + rightProfit[i+1]
		}
	}
	return result
}

// @lc code=end

