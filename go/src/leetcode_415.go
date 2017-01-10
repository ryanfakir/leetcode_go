package main

import "strconv"
import "strings"

func addStrings(nums1 string, nums2 string) string {
	var carry int
	var temp []string
	len1, len2 := len(nums1), len(nums2)
	for len1 > 0 || len2 > 0 || carry == 1 {
		var i, j int
		if len1 <= 0 {
			i = 0
		} else {
			i, _ = strconv.Atoi(string(nums1[len1-1]))
		}
		if len2 <= 0 {
			j = 0
		} else {
			j, _ = strconv.Atoi(string(nums2[len2-1]))
		}
		sum := i + j + carry
		temp = append(temp, strconv.Itoa(sum%10))
		carry = sum / 10
		len1--
		len2--
	}
	for left, right := 0, len(temp)-1; left < right; left, right = left+1, right-1 {
		temp[left], temp[right] = temp[right], temp[left]
	}
	return strings.Join(temp, "")
}
