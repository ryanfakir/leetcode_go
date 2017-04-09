package main

import "strings"

func findLongestWord(s string, d []string) string {
	var res string
	for _, v := range d {
		var i int
		for _, el := range s {
			if len(v) == i {
				break
			}
			if rune(v[i]) == el {
				i++
			}
		}
		if i == len(v) && len(v) >= len(res) {
			if i > len(res) || strings.Compare(v, res) == -1 {
				res = v
			}
		}
	}
	return res
}
