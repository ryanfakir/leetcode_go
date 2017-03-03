package main

import "strconv"

func readBinaryWatch(num int) []string {
	hours := []int{1, 2, 4, 8}
	mins := []int{1, 2, 4, 8, 16, 32}
	var res []string
	for i := 0; i <= num; i++ {
		hourArr := generate(i, hours)
		minArr := generate(num-i, mins)
		for _, v := range hourArr {
			if v >= 12 {
				continue
			}
			for _, el := range minArr {
				if el >= 60 {
					continue
				}
				var s = "" + strconv.Itoa(v) + ":"
				if el > 9 {
					s += strconv.Itoa(el)
				} else {
					s += "0" + strconv.Itoa(el)
				}
				res = append(res, s)
			}
		}
	}
	return res
}

func generate(nums int, pool []int) []int {
	return helper(nil, 0, nums, 0, pool)
}

func helper(res []int, total, size, level int, pool []int) []int {
	if size == 0 {
		res = append(res, total)
		return res
	}
	for i := level; i < len(pool); i++ {
		res = helper(res, total+pool[i], size-1, i+1, pool)
	}
	return res
}
