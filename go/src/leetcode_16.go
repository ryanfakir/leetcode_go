func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    res := 1 << 8
    var fSum int
    for i:= 0; i < len(nums)-2; i++ {
        if i - 1>= 0 && nums[i-1] == nums[i] {continue}
        j , k := i+1, len(nums)-1
        for j < k {
            sum := nums[i] + nums[j] + nums[k] - target
            if math.Abs(float64(sum)) - float64(res) < 0 {
                res = int(math.Abs(float64(sum)))
                fSum = nums[i] + nums[j] + nums[k]
            }
            if sum == 0 {
                return target
            }
            if sum < 0 {
                j++
            } 
            if sum > 0 {
                k--
            }
            
        }
    }
    return fSum
}
