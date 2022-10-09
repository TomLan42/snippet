/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {

	front := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			swap(nums, front, i)
			front++
		}
	}

	return front
}

func swap(nums []int, x, y int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}

// @lc code=end

