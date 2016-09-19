package src

import "math"
func titleToNumber(s string) int {
    var res int
    for i, e := range s {
        res  += int(math.Pow(26, float64(len(s)- i - 1)))*int((e-'A'+1))
    }
    return res
}