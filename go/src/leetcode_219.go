package src

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int, 0)
	for i, el := range nums {
		if num, ok := dict[el]; ok {
			if i-num <= k {
				return true
			}
		}
		dict[el] = i
	}
	return false
}
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if m[v] != 0 && i-m[v]+1 <= k {
			return true
		}
		m[v] = i + 1

	}
	return false
}
