package main

import (
	"bytes"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Sort(ArrayEl(nums))
	var buff bytes.Buffer
	for _, v := range nums {
		buff.WriteString(strconv.Itoa(v))
	}
	res := buff.String()
	if len(res) > 0 {
		if res[0] == '0' {
			return "0"
		}
	}
	return res
}

type ArrayEl []int

func (a ArrayEl) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ArrayEl) Less(i, j int) bool {
	return strconv.Itoa(a[i])+strconv.Itoa(a[j]) > strconv.Itoa(a[j])+strconv.Itoa(a[i])
}
func (a ArrayEl) Len() int { return len(a) }
