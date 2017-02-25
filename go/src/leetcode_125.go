package main

import "strings"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if !isLetter(s[i]) {
			i++
			continue
		} else if !isLetter(s[j]) {
			j--
			continue
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		j--
		i++
	}
	return true
}

func isLetter(letter byte) bool {
	if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') || (letter >= '0' && letter <= '9') {
		return true
	}
	return false
}
