package src

func canPartition(nums []int) bool {
    len := len(nums)
    sum := 0
    for _, el := range nums {
        sum += el
    }
    if sum %2 != 0 {
        return false
    }
    dp := make([][]bool, sum/2+1)
    for i, _ := range dp {
        dp[i] = make([]bool, len+1)
    }

    for i := 0; i<=len;i++ {
        dp[0][i] = true
    }

    for i:= 1; i<=sum/2; i++{
        dp[i][0] = false
    }

    for i:=1; i<=len; i++ {
        for j:= 1; j<=sum/2; j++ {
            dp[j][i] = dp[j][i-1]
            if j >= nums[i-1] {
               dp[j][i] = dp[j][i] || dp[j-nums[i-1]][i-1]
            }
        }
    }
    return dp[sum/2][len]
}
