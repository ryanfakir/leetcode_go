func bulbSwitch(n int) int {
    var res int
    var x int = 1
    for x * x <= n {
        res++
        x++
    }
    return res
}
