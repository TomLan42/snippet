package main

func MergeSort(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	left := MergeSort(nums[:len(nums)/2])
	right := MergeSort(nums[len(nums)/2:])

	result := make([]int, len(nums))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, l)
			l++
		} else {
			result = append(result, r)
			r++
		}
	}

	if l < len(left) {
		result = append(result, left[l:]...)
	}

	if r < len(right) {
		result = append(result, right[r:]...)
	}

	return result
}

func QuickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, 0, len(nums))
		QuickSort(nums, start, pivot-1)
		QuickSort(nums, pivot+1, end)
	}

}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	front := nums[start]
	for i := start; i < end; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, front, i)
			front++
		}
	}

	swap(nums, end, front)
	return front
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func findMedian(nums1, nums2 []int) int {

	// make sure nums1 is the longer array
	if len(nums2) > len(nums1) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}

	// [x x ] [x x x]
	// [- -][-]

	total := len(nums1) + len(nums2)
	left, right := 0, len(nums1)-1

	for left < right {
		mid1 := left + (right-left)/2
		mid2 := total - mid1

		if nums1[mid1] > nums2[mid2+1] {
			right = mid1
		}

		if nums1[mid1+1] < nums2[mid2] {
			left = mid1
		}
	}

}
