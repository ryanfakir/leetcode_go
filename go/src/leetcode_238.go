package src

func productExceptSelf(nums []int) []int {
    length := len(nums)
    res := make([]int, length)
    res[0] = 1

    for i:=1; i<length; i++ {
        res[i] = res[i-1]*nums[i-1]
    }
    right := 1
    for i := length-1; i>=0;i-- {
        res[i] *= right
        right *= nums[i]
    }
    return res
}
