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
