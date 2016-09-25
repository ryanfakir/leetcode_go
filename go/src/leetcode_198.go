package src

import "math"
func rob(nums []int) int {
    if (len(nums) == 0) {
        return 0
    }
    s := make([]int, len(nums)+1)
    s[0], s[1] = 0, nums[0]
    for i, _ := range s {
        if i < 2 {
            continue
        }
        s[i] = int(math.Max(float64(s[i-1]), float64(s[i-2]+ nums[i-1])))
    }
    return s[len(nums)]
}
