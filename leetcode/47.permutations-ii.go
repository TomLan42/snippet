/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	return helper(nums)
}

func helper(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{nums}
	}

	results := make([][]int, 0)

	ashead := make(map[int]bool)
	for i := 0; i < len(nums); i++ {

		if ashead[nums[i]] {
			continue
		}
		ashead[nums[i]] = true
		swap(0, i, nums)
		lists := helper(nums[1:])
		for j := 0; j < len(lists); j++ {
			result := []int{nums[0]}
			result = append(result, lists[j]...)
			results = append(results, result)
		}
		swap(0, i, nums)
	}
	return results
}

func swap(i, j int, nums []int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// @lc code=end

