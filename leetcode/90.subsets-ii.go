/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	cur := make([]int, 0)
	result := make([][]int, 0)
	nums = mergeSort(nums)
	backtrack(nums, 0, cur, &result)
	return result
}

func backtrack(nums []int, pos int, cur []int, result *[][]int) {
	temp := make([]int, len(cur))
	copy(temp, cur)
	*result = append(*result, temp)

	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		backtrack(nums, i+1, cur, result)
		cur = cur[0 : len(cur)-1]
	}
	return
}

func mergeSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	l1 := mergeSort(nums[:mid])
	l2 := mergeSort(nums[mid:])
	return merge(l1, l2)
}

func merge(left, right []int) []int {
	result := make([]int, 0)

	l := 0
	r := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result

}

// @lc code=end

