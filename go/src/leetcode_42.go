package main

import "math"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	var res int
	for left < right {
		min := int(math.Min(float64(height[left]), float64(height[right])))
		if min == height[left] {
			left++
			for left < right && height[left] < min {
				res += min - height[left]
				left++
			}
		} else {
			right--
			for left < right && height[right] < min {
				res += min - height[right]
				right--
			}
		}
	}
	return res
}


func trap(height []int) int {
    x , y := 0, len(height)-1
    var lMax, rMax, max int
    for x <= y {
        lMax = int(math.Max(float64(lMax), float64(height[x])))
        rMax = int(math.Max(float64(rMax), float64(height[y])))
        if lMax > rMax {
            max += rMax - height[y]
            y--
        } else {
            max +=  lMax - height[x]
            x++
        }
    }
    return max
}
