package main

func readBinaryWatch(num int) []string {
	hours := []int{1, 2, 4, 8}
	mins := []int{1, 2, 4, 8, 16, 32}
	var hoursArr, minsArr [][]int
	for i := 1; i < nums; i++ {
		hoursArr = append(hoursArr, generate(i)...)
		minsArr = append(minsArr, generate(nums-i)...)
	}
	for _, v := range hoursArr {

	}

}

func generate(nums int, len int) [][]int {
	return helper(nil, nil, nums, 0, len)
}

func helper(res [][]int, temp []int, nums int, level int, len int) [][]int {
	if len(temp) == nums {
		newtemp := make([]int, len(temp))
		copy(newtemp, temp)
		res = append(res, newtemp)
		return res
	}
	for i := level; i < len; i++ {
		temp = append(temp, level)
		res = helper(res, temp, nums, i+1, len)
		temp = temp[:len(temp)-1]
	}
	return res
}
