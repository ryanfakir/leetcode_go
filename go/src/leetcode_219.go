package src

func containsNearbyDuplicate(nums []int, k int) bool {
    dict := make(map[int]int, 0)
    for i, el := range nums {
        if num, ok := dict[el]; ok {
            if i-num <= k {
                return true
            }
        }
        dict[el] = i
    }
    return false
}
