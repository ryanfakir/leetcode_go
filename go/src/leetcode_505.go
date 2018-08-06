func shortestDistance(maze [][]int, start []int, destination []int) int {
    row, col := len(maze), len(maze[0])
    dp := make([][]int, row)
    for i := range dp {
        dp[i] = make([]int, col)
    }
    for i := 0; i< row; i++ {
        for j := 0; j< col; j++ {
            dp[i][j] = 1 << 31-1
        }
    }
    dp[start[0]][start[1]] = 0
    dfs(maze, dp, start, destination)
    if dp[destination[0]][destination[1]] == 1 << 31 -1 {
        return -1
    }
    return dp[destination[0]][destination[1]]
}


func dfs(maze, dp [][]int, start, end []int)  {
    dirs := [][]int{{0,1}, {0, -1}, {1, 0}, {-1, 0}}
    for _ , dir := range dirs {
        var cnt int
        x, y := start[0], start[1] 
        for x >=0 && y >=0 && x < len(maze) && y < len(maze[0]) && maze[x][y] != 1 {
            x +=dir[0]
            y += dir[1]
            cnt++
        }
        x -= dir[0]
        y -= dir[1]
        cnt--
        if dp[start[0]][start[1]] + cnt < dp[x][y] {
            dp[x][y] = dp[start[0]][start[1]] + cnt
            dfs(maze, dp, []int{x, y}, end)
        }
    }
}
