package src

import (
	"strconv"
)
func addBinary(a string, b string) string {
    if len(a) < len(b) {
        temp := b
        b = a
        a = temp
    }
    lenA := len(a)
    lenB := len(b)
    carry := 0
    digit := ""
    for i:=lenB-1;i>=0;i-- {
        rb, _:= strconv.Atoi(b[i:i+1])
        ra, _ := strconv.Atoi(a[i+lenA-lenB:i+lenA-lenB+1])
        sum := ra + rb + carry
        carry = sum / 2
        digit += strconv.Itoa(sum%2)
    }
    for i :=lenA-lenB-1; i>=0; i-- {
        ra, _:= strconv.Atoi(a[i:i+1])
        sum := ra + carry
        carry = sum/2
        digit += strconv.Itoa(sum%2)
    }
    if carry == 1 {
        digit += "1"
    }
    // reverse
    runes := []rune(digit)
    for i,j := 0, len(runes)-1; i < j; i,j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
