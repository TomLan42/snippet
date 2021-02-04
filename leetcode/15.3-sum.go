/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	// sort the array
	heapSort(nums)

	// mid pointer + left/right pointers
	mid := 0
	left, right := 0, 0
	length := len(nums)
	result := make([][]int, 0)

	for mid = 1; mid < length-1; mid++ {
		left, right = 0, length-1

		if mid > 1 && nums[mid] == nums[mid-1] {
			left = mid - 1
		}

		for left < mid && right > mid {
			// check dup
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < length-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			sum := nums[mid] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[left], nums[mid], nums[right]})
				left++
				right--
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}

	}

	return result

}

func heapSort(nums []int) {

	// construct heap
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i, len(nums))
	}

	for i := len(nums) - 1; i >= 1; i-- {
		swap(nums, 0, i)
		sink(nums, 0, i)
	}

}

func sink(nums []int, i int, length int) {

	for {
		l := 2*i + 1
		r := 2*i + 2

		//idx to swap with
		idx := i

		if l < length && nums[l] > nums[idx] {
			idx = l
		}
		if r < length && nums[r] > nums[idx] {
			idx = r
		}
		if idx == i {
			break
		}
		swap(nums, i, idx)
		i = idx
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// @lc code=end

