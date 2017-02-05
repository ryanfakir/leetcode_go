package src

func removeElement(nums []int, val int) int {
	b := nums[:0]
	for _, e := range nums {
		if e != val {
			b = append(b, e)
		}
	}
	return len(b)
}
func removeElement(nums []int, val int) int {
	var res int
	for _, v := range nums {
		if v != val {
			nums[res] = v
			res++
		}
	}
	return res
}
