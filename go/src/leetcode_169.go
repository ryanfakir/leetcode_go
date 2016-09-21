package src

import "sort"
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[(len(nums)-1)/2]
}