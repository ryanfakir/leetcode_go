package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, numCourses)
	m := make(map[int][]int)
	for _, v := range prerequisites {
		m[v[1]] = append(m[v[1]], v[0])
		degree[v[0]]++
	}
	var q []int
	for i, v := range degree {
		if v == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		for _, v := range m[pop] {
			degree[v]--
			if degree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	for _, v := range degree {
		if v != 0 {
			return false
		}
	}
	return true
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1])
	}
	visited := make(map[int]bool)
	stack := make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if dfs(i, visited, stack, graph) {
			return false
		}
	}
	return true
}

func dfs(node int, visited, stack map[int]bool, graph map[int][]int) bool {
	if !visited[node] {
		visited[node] = true
		stack[node] = true
		for _, v := range graph[node] {
			if !visited[v] && dfs(v, visited, stack, graph) {
				return true
			} else if stack[v] {
				return true
			}
		}
	}
	stack[node] = false
	return false
}


func canFinish(numCourses int, prerequisites [][]int) bool {
    pre := prerequisites
    m := make(map[int][]int)
    for _, v := range pre {
        m[v[0]] = append(m[v[0]], v[1])
    }
    visited:= make(map[int]int)
    for i:=0 ; i < numCourses; i++ {
        if  !dfs(visited, i, m) {return false}
    }
    return true
}


func dfs(visited map[int]int, i int, grid map[int][]int) bool {
    if visited[i] == -1 {return false}
    if visited[i] == 1 {return true}
    visited[i] = -1
    for _, v := range grid[i] {
        if !dfs(visited, v, grid) {return false}
    }
    visited[i]= 1
    return true
}
