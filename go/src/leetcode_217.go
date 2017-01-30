package src

func containsDuplicate(nums []int) (res bool) {
	hm := make(map[int]int)
	res = false
	for _, e := range nums {
		if _, ok := hm[e]; ok {
			res = true
		} else {
			hm[e] = 0
		}
	}
	return
}
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
