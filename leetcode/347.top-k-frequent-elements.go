/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	bucket := make([][]int, len(nums)+1)

	for k, v := range count {
		bucket[v] = append(bucket[v], k)
	}

	top := k
	result := []int{}
	for i := len(nums); i >= 0; i-- {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
			top--
		}
		if len(result) >= k {
			break
		}
	}
	return result
}

// @lc code=end
