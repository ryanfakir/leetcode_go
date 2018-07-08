func evalRPN(tokens []string) int {
    if len(tokens) == 0 {return 0}
    var s []int
    for _, v := range tokens {
        if v != "/" && v != "+" && v != "-" && v != "*" {
            num, _ := strconv.Atoi(v)
            s = append(s, num)
        } else {
            m,n := s[len(s)-2], s[len(s)-1]
            s = s[:len(s)-2]
            var res int
            if v == "/" {
                res = m / n
            } 
            if v == "*" {
                res = m * n
            } 
            if v == "+" {
                res = m + n
            }
            if v == "-" {
                res = m - n
            }
            s = append(s, res)
        }
    }
    return s[0]
}
