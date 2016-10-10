package src

func missingNumber(nums []int) int {
    len := len(nums)
    res := 0;
    for i, el := range nums {
        res ^= i ^ el
    }
    res ^= len
    return res
}
