package src

import "strconv"
func findNthDigit(n int) int {
    digit, base, start := 1,9,1
    for n-base*digit > 0 {
        n -= base*digit
        digit++
        base *=10
        start *=10
    }
    start += (n-1)/digit
    res := strconv.Itoa(start)
    return int(res[(n-1)%digit] - '0')
}