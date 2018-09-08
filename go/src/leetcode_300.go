package main

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				max = int(math.Max(float64(max), float64(dp[j]+1)))
			}
		}
		dp[i] = max
	}
	var res int
	for _, v := range dp {
		res = int(math.Max(float64(res), float64(v)))
	}
	return res
}

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {return 0}
    var tmp []int = []int{nums[0]}
    for _, v := range nums {
        if v <= tmp[0] {
            tmp[0] = v
        } else if v > tmp[len(tmp)-1] {
            tmp = append(tmp, v)
        } else {
            left, right := 0, len(tmp)-1
            for left <= right {
                mid := (left + right) /2
                if tmp[mid] < v {
                    left = mid +1
                } else if tmp[mid] >= v {
                    right = mid-1
                }
            }
            tmp[left] = v
        }
    }
    return len(tmp)
}

func lengthOfLIS(nums []int) int {
    var tmp []int
    for _, v := range nums {
        if len(tmp) == 0 || v > tmp[len(tmp)-1] {
            tmp = append(tmp, v)
        } else {
            l, r := 0, len(tmp)-1
            for l < r {
                mid := (l + r) /2
                if tmp[mid] >= v {
                    r = mid
                } else {
                    l = mid +1
                }
            }
            tmp[l] = v
        }
    }
    return len(tmp)
    
}
