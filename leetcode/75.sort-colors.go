/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {

	red, white, blue := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[blue] = 2
			nums[white] = 1
			nums[red] = 0
			red++
			white++
			blue++
		} else if nums[i] == 1 {
			nums[blue] = 2
			nums[white] = 1
			white++
			blue++
		} else if nums[i] == 2 {
			blue++
		}

	}

}

// @lc code=end

