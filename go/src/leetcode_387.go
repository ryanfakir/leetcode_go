package src

func firstUniqChar(s string) (res int) {
	dict := make([]int, 26)
	for _, e := range s {
		dict[e-'a']++
	}
	res = -1
	for i, e := range s {
		if dict[e-'a'] == 1 {
			res = i
			break
		}
	}
	return
}
