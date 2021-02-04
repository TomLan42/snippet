/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	// return twoPass(nums, target)
	return onePass(nums, target)
}

// twoPass solution
func twoPass(nums []int, target int) []int {

	//init map
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	// go through, 自带去重
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if m[complement] != i && m[complement] > 0 {
			result[0], result[1] = i, m[complement]
		}
	}

	return result

}

// onePass solution
func onePass(nums []int, target int) []int {

	m := make(map[int]int)
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if idx, found := m[complement]; found {
			result[0], result[1] = idx, i
			return result
		} else {
			m[nums[i]] = i
		}
	}
	return result

}

// @lc code=end

