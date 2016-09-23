package src

func isHappy(n int) (res bool) {
    var sum int
    set := make(map[int]bool)
    for n > 0 {
        remain := n%10
        sum += remain * remain
        n= n/10
        if n == 0 {
            if sum == 1 {
                res = true
                break
            }
            if _, ok := set[sum]; ok {
                res = false
                break
            }
            n = sum
            set[sum] = true
            sum = 0
        }
    }
    return
}
