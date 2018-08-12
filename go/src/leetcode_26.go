package src

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    b := nums[:1]
    for i, _ := range nums {
        if i== 0 {
            continue
        }
        if nums[i] != nums[i-1] {
            b = append(b, nums[i])
        }
    }
    return len(b)
}

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {return 0}
    prev := nums[0]
    var res int
    for i:= 1; i < len(nums); i++ {
        if nums[i] != prev {
            res++
            nums[res] = nums[i]
            prev = nums[res]
        }
    }
    return res + 1
}
