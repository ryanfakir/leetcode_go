package src

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    return x == reverse(x)
}

func reverse(x int) int {
    res := 0
    for x>0 {
        res = res*10 + x%10
        x /= 10
    }
    return res
}