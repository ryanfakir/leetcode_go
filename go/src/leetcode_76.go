package main

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	dict := make(map[string]int)
	for _, v := range t {
		dict[string(v)]++
	}
	var left, right, count int
	var res = s + "string"
	var length = len(res)
	for ; right < len(s); right++ {
		if _, ok := dict[string(s[right])]; ok {
			dict[string(s[right])]--
			if dict[string(s[right])] >= 0 {
				count++
			}
			for count == len(t) {
				if right-left+1 <= len(res) {
					res = s[left : right+1]
				}
				if _, ok := dict[string(s[left])]; ok {
					dict[string(s[left])]++
					if dict[string(s[left])] > 0 {
						count--
					}
				}
				left++
			}
		}
	}
	if length == len(res) {
		res = ""
	}
	return res
}
