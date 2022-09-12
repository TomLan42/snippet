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

func QuickSort(nums []int) []int {

	k := 

}

func partition(nums []int, start, end int) {

	piviot := nums[end]
	front := start
	for i:= start; i < end;i++{

	}

}

func swap(nums []int, i, j int)
