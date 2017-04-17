package main

func removeInvalidParentheses(s string) []string {
	var q []string
	q = append(q, s)
	var res []string
	var rightLevel bool
	visited := make(map[string]bool)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]
			if valid(pop) {
				rightLevel = true
				res = append(res, pop)
				continue
			}
			for i := 0; i < len(pop); i++ {
				if pop[i] != '(' && pop[i] != ')' {
					continue
				}
				temp := pop[:i] + pop[i+1:]
				if !visited[temp] {
					visited[temp] = true
					q = append(q, temp)
				}
			}
		}
		if rightLevel {
			break
		}
	}
	return res
}

func valid(s string) bool {
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		}
		if s[i] == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}
