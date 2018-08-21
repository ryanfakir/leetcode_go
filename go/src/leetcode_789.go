func escapeGhosts(ghosts [][]int, target []int) bool {
    dist := math.Abs(float64(target[0])) + math.Abs(float64(target[1]))
    mn := float64(1 << 15 -1)
    for _, v := range ghosts {
        t := math.Abs(float64(v[0] - target[0])) + math.Abs(float64(v[1] - target[1]))
        mn = math.Min(t, mn)
    }
    return dist < mn
}
