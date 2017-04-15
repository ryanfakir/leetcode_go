package src

import (
	"bytes"
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
	for i := lenB - 1; i >= 0; i-- {
		rb, _ := strconv.Atoi(b[i : i+1])
		ra, _ := strconv.Atoi(a[i+lenA-lenB : i+lenA-lenB+1])
		sum := ra + rb + carry
		carry = sum / 2
		digit += strconv.Itoa(sum % 2)
	}
	for i := lenA - lenB - 1; i >= 0; i-- {
		ra, _ := strconv.Atoi(a[i : i+1])
		sum := ra + carry
		carry = sum / 2
		digit += strconv.Itoa(sum % 2)
	}
	if carry == 1 {
		digit += "1"
	}
	// reverse
	runes := []rune(digit)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		temp := a
		a = b
		b = temp
	}
	al, bl := len(a), len(b)
	al--
	bl--
	var carry int
	var buffer bytes.Buffer
	for al >= 0 && bl >= 0 {
		an, _ := strconv.Atoi(string(a[al]))
		bn, _ := strconv.Atoi(string(b[bl]))
		buffer.WriteString(strconv.Itoa((an + bn + carry) % 2))
		carry = (an + bn + carry) / 2
		al--
		bl--
	}
	for al >= 0 {
		an, _ := strconv.Atoi(string(a[al]))
		buffer.WriteString(strconv.Itoa((an + carry) % 2))
		carry = (an + carry) / 2
		al--
	}
	if carry == 1 {
		buffer.WriteString("1")
	}
	arr := buffer.Bytes()
	if len(arr) > 1 {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return string(arr)
}
