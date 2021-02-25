## Edit distance

'abcd'
'dfef'

``` go
func editSpace(str1 string, str2 string) int {

    // init dp array
    dp := make([][]int, len(str1)+1)
    for i := 0; i < len(dp); i++{
        dp[i] = make([]int,len(str2)+1)
    }

    for i := 0; i < len(dp); i++ {
        dp[i][0] = i
    }

    for j := 0; j < len(dp[0]); j++ {
        dp[0][j] = j
    }

    // state transfer
    for i := 1; i < len(dp); i++ {
        for j:=1; j < len(dp[i]); j++ {
            
            if str1[i-1] == str[i-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1])
            }
        }
    }
    return dp[len(str1)-1][len(str2)-1]
}
```