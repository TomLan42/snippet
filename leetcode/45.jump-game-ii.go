/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {

	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		idx := 0
		for idx < len(nums) && idx+nums[idx] < i {
			idx++
		}
		f[i] = f[idx] + 1
	}

	return f[len(nums)-1]

}

// @lc code=end

