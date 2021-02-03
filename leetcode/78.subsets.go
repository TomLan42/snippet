/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// final result
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result

}

func backtrack(nums []int, pos int, list []int, result *[][]int) {

	temp := make([]int, len(list))
	copy(temp, list)
	*result = append(*result, temp)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
	return
}

// @lc code=end

