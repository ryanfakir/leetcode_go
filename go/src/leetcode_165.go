package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	var i, j int
	for i < len(v1) || j < len(v2) {
		var a1, a2 int
		if i >= len(v1) {
			a1 = 0
		} else {
			a1, _ = strconv.Atoi(v1[i])
		}
		if j >= len(v2) {
			a2 = 0
		} else {
			a2, _ = strconv.Atoi(v2[j])
		}
		if a1 > a2 {
			return 1
		} else if a1 < a2 {
			return -1
		}
		i++
		j++
	}
	return 0
}
