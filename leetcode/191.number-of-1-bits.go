/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {

	mask := uint32(1)

	counter := 0
	for i := 0; i < 32; i++ {
		result := mask & (num >> i)
		if result == 1 {
			counter++
		}
	}
	return counter

}

// @lc code=end

