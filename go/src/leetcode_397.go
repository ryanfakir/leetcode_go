package src

import "math"
func integerReplacement(n int) int {
    if n == 1 {
        return 0
    }
    if n & 1 == 0 {
        return integerReplacement(n/2) +1
    }
    return int(math.Min(float64(integerReplacement(n-1) + 1), float64(integerReplacement(n+1) + 1)))
}
