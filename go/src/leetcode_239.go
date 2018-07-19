func maxSlidingWindow(nums []int, k int) []int {
    var res []int
    var q []int
    for i:=0; i < len(nums);i++ {
        if len(q) > 0 && q[0] <= i - k { q= q[1:]}
        for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {q = q[:len(q)-1]}
        q = append(q, i)
        if i - k >= -1 {
            res = append(res, nums[q[0]])
        }
    }
    return res
}
