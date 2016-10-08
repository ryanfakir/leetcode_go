package src

func singleNumber(nums []int) []int {
    xor := 0
    for _, el := range nums{
        xor ^= el
    }
    lastBit := xor - (xor&(xor-1))
    groupA, groupB := 0, 0
    for _,el := range nums {
        if (el&lastBit) == 0 {
            groupA ^= el
        } else {
            groupB ^= el
        }
    }
    res := make([]int, 0)
    res = append(res, groupA)
    res = append(res, groupB)
    return res
}
