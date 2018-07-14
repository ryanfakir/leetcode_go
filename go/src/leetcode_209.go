func minSubArrayLen(s int, nums []int) int {
    var left, right int
    var sum int
    res := len(nums) + 1
    for right < len(nums)  {
        for sum < s && right < len(nums) {
            sum += nums[right]
            right++
        }
        for sum >= s && left < right {
            res = int(math.Min(float64(res), float64(right- left )))
            sum -= nums[left]
            left++
        }
    }
    if res == len(nums) + 1 {
        return 0
    } else {
        return res
    }
}
