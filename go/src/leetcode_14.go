package src

import "unicode/utf8"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, el := range strs {
		if len(res) < len(el) {
			el = el[:len(res)]
		}
		for j, char := range el {
			r, _ := utf8.DecodeRuneInString(res[j : j+1])
			if char != r {
				res = res[:j]
				break
			}
		}
		if len(el) <= len(res) {
			res = el
		}
	}
	
	return res
}



