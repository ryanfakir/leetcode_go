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

func minWindow(s string, t string) string {
    dict := make(map[rune]int)
    for _ , v := range t {
        dict[v]++
    }
    res := s + "s"
    var cnt, left int
    for i := 0; i < len(s); i++ {
        dict[rune(s[i])]--
        if dict[rune(s[i])] >= 0 {
            cnt++
        }
        for cnt == len(t) {
            if len(s[left:i+1]) < len(res) {
                res = s[left: i+1]
            }
            if _, ok := dict[rune(s[left])]; ok {
                dict[rune(s[left])]++
                if dict[rune(s[left])] > 0 {
                    cnt--
                }
            }
            left++
        }
    }
    if len(res) > len(s) {return ""}
    return res
}
