package src

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    state := make([]int, n+1)
    state[1], state[2] = 1,2
    for i :=3; i<n+1; i++ {
        state[i] = state[i-1] + state[i-2]
    }
    return state[n]
}
