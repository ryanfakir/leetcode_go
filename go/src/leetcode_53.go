func maxSubArray(nums []int) int {
    var res = nums[0]
    local := nums[0]
    for i:= 1 ; i < len(nums); i++ {
            if  nums[i] > local + nums[i] {
                local = nums[i]
            } else {
                local += nums[i]
            }
            res = int(math.Max(float64(res), float64(local)))
    }
    return res
}
