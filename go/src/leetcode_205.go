package src

func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte, 0)
	tMap := make(map[byte]byte, 0)
	for i := 0; i < len(s); i++ {
		if e, ok := sMap[s[i]]; ok {
			if e != t[i] {
				return false
			}
		} else {
			sMap[s[i]] = t[i]
		}

		if e, ok := tMap[t[i]]; ok {
			if e != s[i] {
				return false
			}
		} else {
			tMap[t[i]] = s[i]
		}
	}
	return true
}
func isIsomorphic(s string, t string) bool {
	m1, m2 := make([]int, 256), make([]int, 256)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}
