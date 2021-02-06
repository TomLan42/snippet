/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return heapMethod(nums, k)
}

// the heap method
func heapMethod(nums []int, k int) int {

	// construct min heap of the frist k elements
	for i := k/2 - 1; i >= 0; i-- {
		sink(i, nums[:k], k)
	}

	// keep insert the rest
	for i := k; i < len(nums); i++ {
		if nums[i] < nums[0] {
			continue
		}
		swap(i, 0, nums)
		sink(0, nums[:k], k)
	}

	return nums[0]

}

func sink(i int, nums []int, length int) {
	for {
		l := 2*i + 1
		r := 2*i + 2

		idx := i

		if l < length && nums[l] < nums[idx] {
			idx = l
		}
		if r < length && nums[r] < nums[idx] {
			idx = r
		}
		if idx == i {
			break
		}
		swap(i, idx, nums)
		i = idx
	}
}

func swap(i, j int, nums []int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// @lc code=end

