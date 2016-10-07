package src

import "math"
func reverse(x int) int {
    n := x
    positive := false
    if n > 0 {
        positive = true
    }
    res := 0
    for n != 0 {
        n = int(math.Abs(float64(n)))
        digit := n%10
        n /= 10
        res = res*10 + digit
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    if positive {
        return res
    } else {
        return -res
    }
}
