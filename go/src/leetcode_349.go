package src

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	if nums1 == nil || nums2 == nil {
		return make([]int, 0)
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1Len, nums2Len := len(nums1), len(nums2)
	res := make([]int, 0)
	for i, j := 0, 0; i < nums1Len && j < nums2Len; {
		if nums1[i] == nums2[j] {
			if len(res) == 0 || res[len(res)-1] != nums1[i] {
				res = append(res, nums1[i])
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
