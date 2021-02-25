## Binary Search Template


``` go

func search(nums []int, target int) int {

    left := 0
    right := len(nums)-1

    mid := left + (right-left)/2

    for left < right {
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            right = mid
        }
        if nums[mid] < target {
            left = mid
        }
    }

    if nums[left] == target {
        return left
    }
    if nums[right] == target {
        return right
    }
    return -1
}


```


## Binary Search Range
在binary search之上对区间变化有更subtle的控制

```go

func searchRange(nums[]int) []int{

    result := []int{-1,-10}

    left := 0 
    right := len(nums)-1

    mid := left + (right-left)/2

    // find the left border
    for left < right {
        if nums[mid] > target {
            right = mid
        }
        if nums[mid] < target {
            left = mid
        }
        if nums[mid] == target {
            right = mid
        }
    }
    if nums[right] == target {
        result[0] = right
    }
    if nums[left] == target {
        result[0] = left
    }


    for left < right {
        if nums[mid] > target {
            right = mid
        }
        if nums[mid] < target {
            left = mid
        }
        if nums[mid] == target {
            left = mid
        }
    }

    if nums[left] == target {
        result[1] = left
    }
    if nums[right] == target {
        result[1] = right
    }

    return result

}

```