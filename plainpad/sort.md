## 归并排序 Merge Sort

```go


func MergeSort(nums []int) []int {

    if len(nums) == 1 {
        return nums
    }

    left := mergeSort(nums[:len(nums)/2])
    right := mergeSort(nums[len(nums)/2:])

    idx := 0 
    result := make([]int,0)
    
    l, r := 0, 0
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

```

## 快速排序

```go


func QuickSort(nums []int) []int {
    result := quicksort(nums, 0, len(nums)-1)
    return result
}

func quicksort(nums []int, start, end) {
    if start < end {
        pivot := partition(nums, start, end)
        quicksort(nums, start, pivot-1)
        quicksort(nums, pivot+1, end)
    }

}



func partition(nums []int, start, end) int {

    pivot := nums[end]
    front := start
    for i := start; i < end; i++ {
        if nums[i] < pivot {
            swap(nums, i, front)
            front++
        }
    }

    swap(nums, end, front)
    return front
}

func swap(nums []int, i,j int) {
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}


```

## HeapSort 

``` go 

func HeapSort(nums []int) []int {

    // construct max heap
    for i := len(nums)/2 - 1; i >=0; i-- {
        sink(nums, i)
    }



    // move top to rear
    for i := len(nums)-1; i >= 0; i--{
        swap(nums,0,i)
        sink(nums,0,i)
    } 



}


func sink(nums []int, i int, length int) {

    //left, and right child
    for {
        left := i * 2 + 1
        right := i *2 + 2

        idx := i 

        if nums[right] > nums[i] && right < length {
            idx := right 
        }
        if nums[left] > nums[i] && left < length {
            idx := left
        }

        if idx == i {
            break
        }
        swap(nums, idx, i)
        i = idx
    }

}

func swap(nums []int, i , j int) {
    nums[i] = temp 
    nums[i] = nums[j]
    nums[j] = temp
}
```
