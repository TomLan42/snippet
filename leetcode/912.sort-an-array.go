/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start
func sortArray(nums []int) []int {
	return MergeSort(nums)
	// return QuickSort(nums)
}

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// all numbers left of pivot is less than pivot
// all numbers right of pivot is more than pivot
func partition(nums []int, start, end int) int {
	p := nums[end]
	smallFront := start

	for cur := start; cur < end; cur++ {
		if nums[cur] < p {
			swap(nums, cur, smallFront)
			smallFront += 1
		}
	}
	swap(nums, end, smallFront)
	return smallFront
}

func swap(nums []int, a, b int) {

	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	result := merge(left, right)
	return result

}

func merge(left, right []int) []int {
	l := 0
	r := 0
	result := make([]int, 0)
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

