func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {return false}
    var small, big = make([]int, len(nums)), make([]int, len(nums))
    small[0] = nums[0]
    big[len(nums)-1] = nums[len(nums)-1]
    for i := 1; i < len(nums); i++ {
        small[i] = int(math.Min(float64(small[i-1]), float64(nums[i])))
    }
    for i := len(nums)-2; i>=0; i-- {
        big[i] = int(math.Max(float64(big[i+1]), float64(nums[i])))
    }
    for i:=0 ; i < len(nums); i++ {
        if small[i] < nums[i] && nums[i] < big[i] {return true}
    }
    return false
}
