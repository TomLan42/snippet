## subset

``` go


func subset(nums []int) [][]int {

    list := make([]int, 0)
    result := make([][]int, 0)

    backtracking(nums, 0, list, &result)
    return result

}



func backtracking(nums []int, pos int, list []int, result *[][]int) {
    
    ans := make([]int, len(list))
    copy(ans, list)
    result = append(result, ans)

    for i := pos; i < len(nums); i++ {
        list = append(list, nums[i])
        backtracking(nums, i+1, list, result)
        list = list[:len(list)-1]
    }

}


```



## permutation

``` go 


```

