func fractionToDecimal(numerator int, denominator int) string {
    var res string
    num, den:= numerator, denominator
    if num == 0 {return "0"}
    if (num > 0) != (den > 0) {
        res += "-"
    }
    num = int(math.Abs(float64(num)))
    den = int(math.Abs(float64(den)))
    res += strconv.Itoa(num/den)
    rem:= num%den
    if rem == 0 {
        return res
    }
    res += "."
    m := make(map[int]int)
    for rem != 0 {
        if m[rem] != 0 {
            res = res[:m[rem]] + "(" + res[m[rem]:]
            res += ")"
            return res
        } else {
            m[rem] = len(res) 
        }
        rem *= 10
        res += strconv.Itoa(rem/den)
        rem %= den
    }
    return res
}
