package src

import "math"
func maxRotateFunction(A []int) int {
    sumUp := 0
    fResult := 0
    length := len(A)
    for i, el := range A {
        fResult += el*i
        sumUp += el
    }
    res := fResult
    for i:= 1; i < length; i++ {
        fResult += sumUp - length*A[length-i]
        res = int(math.Max(float64(res), float64(fResult)))
    }
    return res
}
