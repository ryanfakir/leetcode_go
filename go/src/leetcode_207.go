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
