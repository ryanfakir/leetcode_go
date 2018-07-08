func maxProduct(nums []int) int {
    if len(nums) == 0 {return 0}
    var max, min, res= nums[0], nums[0], nums[0]
    for i:= 1; i< len(nums);i++ {
        tmax, tmin := max, min
        max = int(math.Max(float64(nums[i] * tmin), math.Max(float64(tmax * nums[i]), float64(nums[i]))))
        min = int(math.Min(float64(nums[i] * tmin), math.Min(float64(tmax * nums[i]), float64(nums[i]))))
        res = int(math.Max(float64(res), float64(max)))
    }
    return res
}
