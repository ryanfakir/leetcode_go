package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int
	graph := make(map[int][]int)
	counter := make([]int, numCourses)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		counter[v[0]]++
	}
	var queue []int
	for i, v := range counter {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		lengh := len(queue)
		for i := 0; i < lengh; i++ {
			pop := queue[0]
			queue = queue[1:]
			res = append(res, pop)
			for _, v := range graph[pop] {
				counter[v]--
				if counter[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
	}
	if len(res) < numCourses {
		return []int{}
	}
	return res
}

var res []int

func findOrder(numCourses int, prerequisites [][]int) []int {
    res = nil
    pre := prerequisites
    m := make(map[int][]int)
    for _, v := range pre {
        m[v[0]] = append(m[v[0]], v[1])
    }
    visited:= make(map[int]int)
    for i:= 0; i< numCourses; i++ {
        if visited[i] == 0 && !dfs(m, i, visited) {
            return nil
        }
    }
    return res
}

func dfs(graph map[int][]int, i int, visited map[int]int) bool {
    if visited[i] == 1 {return true}
    if visited[i] == -1 {return false}
    visited[i] = -1 
    for _, v := range graph[i] {
        if !dfs(graph, v, visited) {return false}
    }
    visited[i]= 1
    res = append(res, i)
    return true
}
