/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {

	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] > len(nums)/2 {
			return num
		}
	}

	return 0

}

// @lc code=end

