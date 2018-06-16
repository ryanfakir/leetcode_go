package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := len(nums) - 1; j > 2; j-- {
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}
			left, right := i+1, j-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[left], nums[right], nums[j]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else if sum > target {
					right--
				}
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
    if len(nums) < 4 {return nil}
    var res [][]int
    sort.Ints(nums)
    for i := 0; i<len(nums)-3; i++ {
        if i-1 >= 0 && nums[i-1] ==nums[i] {continue}
        for j := i+1; j < len(nums)-2; j++ {
            if j > i+1 && nums[j-1] == nums[j] {continue}
            k, l := j+1, len(nums)-1
            for k < l {
                sum := nums[i] + nums[j] + nums[k] + nums[l]
                if sum > target {
                    l--
                } else if sum < target {
                    k++
                } else {
                    res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
                    k++
                    l--
                    for k < l && nums[k] == nums[k-1] {
                        k++
                    }
                    for k < l && nums[l] == nums[l+1] {
                        l--
                    }
                }
            }
        }
    }
    return res
}
