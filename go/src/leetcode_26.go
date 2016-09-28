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
