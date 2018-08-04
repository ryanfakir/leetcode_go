func hasPath(maze [][]int, start []int, destination []int) bool {
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
    visited := make(map[int]bool)
    
    dfs(maze, dp, start, destination, visited, 0)
    fmt.Println(dp)
    if dp[destination[0]][destination[1]] == 1 << 31 -1 {
        return false
    }
    
    return true
}


func dfs(maze, dp [][]int, start, end []int, visited map[int]bool, cnt int)  {
    if cnt < dp[start[0]][start[1]] {
        dp[start[0]][start[1]] = cnt
    }
    dirs := [][]int{{0,1}, {0, -1}, {1, 0}, {-1, 0}}
    for _ , dir := range dirs {
        x, y := start[0], start[1] 
        for x >=0 && y >=0 && x < len(maze) && y < len(maze[0]) && maze[x][y] != 1 {
            x +=dir[0]
            y += dir[1]
            cnt++
        }
        x -= dir[0]
        y -= dir[1]
        cnt--
        if !visited[x*len(maze) + y] {
            visited[x*len(maze) + y] = true
            dfs(maze, dp, []int{x, y}, end, visited, cnt)
        }
    }
}



