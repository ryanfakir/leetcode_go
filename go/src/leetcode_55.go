package main

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

func canJump(nums []int) bool {
    if len(nums) == 1 {return true}
    var bound, n, lMax int
    for i:= 0; i <= bound; i++ {
        lMax = int(math.Max(float64(lMax) , float64(i + nums[i])) )
        if lMax >= len(nums) -1 {return true}
        if i == bound {
            if lMax > bound {
                bound = lMax
                n++
            }
        }
    } 
    return false
}
