func rob(nums []int) int {
    if len(nums) == 0 { return 0}
    if len(nums) == 1 {
        return nums[0]
    }
    return int(math.Max(float64(dp(nums, 0, len(nums)-2)), float64(dp(nums, 1, len(nums)-1))))
}


func dp(nums []int, left, right int) int {
    if right- left <= 0 {return nums[left]}
    arr := make([]int, right +1)
    arr[left] = nums[left]
    arr[left + 1] = int(math.Max(float64(nums[left]), float64(nums[left+1])))
    for i := left+  2 ; i <= right; i++ {
        arr[i] = int(math.Max(float64(arr[i-2] + nums[i]), float64(arr[i-1])))
    }
    return arr[right]
}
