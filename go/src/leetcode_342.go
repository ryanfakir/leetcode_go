package src

func isPowerOfFour(num int) bool {
    if num < 1 {
        return false
    }
    for num % 4 == 0 {
        num >>= 2
    }
    return num == 1
}
