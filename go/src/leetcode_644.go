func findMaxAverage(nums []int, k int) float64 {
    var res, sum float64
    for i, v := range nums {
        if i < k {
            sum += float64(v)
        }
    }
    res = sum / float64(k)
    for i := k ; i< len(nums); i++ {
        sum += float64(nums[i])
        total := sum
        if total > res * float64(i+1) {res = total / float64(i+1)}
        for j := 0; j + k <= i; j++ {
            total -= float64(nums[j])
            if total > res * float64(i-j) {res = total / float64(i-j)}
        }
    }
    return res
}
