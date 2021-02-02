/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {

	return helper(nums)

}

func helper(nums []int) [][]int {

	if len(nums) == 1 {
		return [][]int{nums}
	}

	results := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		swap(0, i, nums)
		lists := helper(nums[1:])
		for j := 0; j < len(lists); j++ {
			result := []int{nums[0]}
			result = append(result, lists[j]...)
			results = append(results, result)
		}
		swap(i, 0, nums)
	}

	return results
}

func swap(i, j int, nums []int) {

	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// @lc code=end

