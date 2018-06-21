package main

import "math"

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var res, next, max, i int
	for max-i >= 0 {
		res++
		for ; i <= max; i++ {
			next = int(math.Max(float64(next), float64(i+nums[i])))
			if next >= len(nums)-1 {
				return res
			}
		}
		max = next
	}
	return 0
}

func jump(nums []int) int {
    if len(nums) == 1 {return 0}
    var bound, n, lMax int
    for i:= 0; i <= bound; i++ {
        lMax = int(math.Max(float64(lMax) , float64(i + nums[i])) )
        if lMax >= len(nums) -1 {return n + 1}
        if i == bound {
            if lMax > bound {
                bound = lMax
                n++
            }
        }
    } 
    return 0
}
