func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {return len(nums)}
    res := 1
    prev := nums[0]
    n := 1
    for i:= 1; i< len(nums); i++ {
        if prev == nums[i]{
            if n <= 1 {
                nums[res] = nums[i]
                n++
                res++
            }
        } else {
            prev = nums[i]
            n = 1
            nums[res] = nums[i]
            res++
        }
    }
    return res
}
