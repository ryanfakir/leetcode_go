func countDigitOne(n int) int {
    var res = 0
    a, b := 1, 1
    for (n > 0) {
        tmp := n%10
        if tmp == 1 {
             res += (n + 8) / 10 * a +  b
        } else {
              res += (n + 8) / 10 * a 
        }
       
        b += tmp * a
        a *= 10
        n /= 10
    }
    return res;
}
