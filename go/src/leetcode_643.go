func findMaxAverage(nums []int, k int) float64 {
    var sum, res float64
    for i, v := range nums {
        if i <= k-1 {
            sum +=float64(v)
        }
    }
    res = sum
    for i:= k ; i< len(nums); i++ {
        sum += float64(nums[i] - nums[i-k])
        res = math.Max(res, sum)
    }
    return res/float64(k)
}
