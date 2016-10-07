package src

import "bytes"
func convertToTitle(n int) string {
    var buffer bytes.Buffer
    for n !=0 {
        n--
        buffer.WriteString(string(rune(n%26 + 'A')))
        n /= 26
    }
    runes := []rune(buffer.String())
    for i, j :=0, len(runes)-1; i<j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
