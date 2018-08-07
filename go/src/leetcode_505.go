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

func shortestDistance(maze [][]int, start []int, destination []int) int {
    row , col := len(maze), len(maze[0])
    var q = [][]int{{start[0],start[1]}}
    dp := make([][]int, row)
    for i := range dp {
        dp[i] = make([]int, col)
    }
    for i:= 0; i < row ; i++ {
        for j:= 0; j< col; j++ {
            dp[i][j] = 1 << 63 -1
        }
    }
    dp[start[0]][start[1]] = 0
    var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for len(q) > 0 {
        pop := q[0]
        q = q[1:]
        x, y := pop[0], pop[1]
        for _ , dir := range dirs {
            i, j, cnt := x, y, dp[x][y]
            for i >=0 && j >=0 && i < row && j < col && maze[i][j] == 0 {
                i += dir[0]
                j += dir[1]
                cnt++
            }
            i -= dir[0]
            j -= dir[1]
            cnt--
            if cnt < dp[i][j] {
                dp[i][j] = cnt 
                q = append(q, []int{i, j})
            }
        }
    }
    if dp[destination[0]][destination[1]] == 1 << 63 -1 {
        return -1
    }
    return dp[destination[0]][destination[1]]
}
