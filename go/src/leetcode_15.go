package main

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := len(nums) - 1; j > 1; j-- {
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}
			k := i + 1
			for k > i && k < j {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[k], nums[j]})
					break
				} else if sum > 0 {
					break
				} else {
					k++
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        if i -1 >= 0 && nums[i-1] == nums[i] {continue} 
        left , right := i+1 , len(nums)-1
        for left < right {
            sum := nums[left] + nums[right] + nums[i]
            if sum > 0 {
                right--
            } else if sum < 0 {
                left++
            } else {
                res = append(res, []int{nums[left], nums[right], nums[i]})
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
            }
        }
    }
    return res
}
