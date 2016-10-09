package src
import "math"
func countNumbersWithUniqueDigits(n int) int {
    n = int(math.Min(float64(10), float64(n)))
    if n == 0 {return 1}
    if n == 1 {return 10}
    temp := 9
    res := 0
    for i:=2; i<=n; i++ {
        res += temp*(11-i)
        temp *= 11-i
    }
    return res + 10
}

