package src


func intersect(nums1 []int, nums2 []int) []int {
    hm := make(map[int]int)
    res := make([]int, 0)
    for _, el := range nums1 {
        hm[el]++
    }
    for _, el := range nums2 {
        if _, ok := hm[el]; ok {
            if hm[el] > 0 {
                res = append(res, el)
                hm[el]--
            }
        }
    }
    return res
}