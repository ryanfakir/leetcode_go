type Item struct {
	s string
	v float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]map[string]float64)
	var res []float64
	// base on equations to generate graph
	for i, v := range equations {
		// check subMap is present or not
		if tempM, ok := m[v[0]]; ok {
			tempM[v[1]] = values[i]
		} else {
			tempM := make(map[string]float64)
			tempM[v[1]] = values[i]
			m[v[0]] = tempM
		}
		if tempm, ok := m[v[1]]; ok {
			tempm[v[0]] = 1.0 / values[i]
		} else {
			tempm := make(map[string]float64)
			tempm[v[0]] = 1.0 / values[i]
			m[v[1]] = tempm
		}
	}
	for k, v := range queries {
		// Any good looking way to handle specific multi-keys check in map
		if _, ok := m[v[0]]; !ok {
			res = append(res, -1.0)
			continue
		}
		if _, ok := m[v[1]]; !ok {
			res = append(res, -1.0)
			continue
		}
		if v[0] == v[1] {
			res = append(res, 1.0)
			continue
		}
		q := []Item{}
		q = append(q, Item{v[0], 1.0})
		visited := make(map[string]bool)
		// BFS Traversal
		for len(q) != 0 {
			var pop Item
			pop, q = q[len(q)-1], q[:len(q)-1]
			if visited[pop.s] {
				continue
			}
			visited[pop.s] = true
			for sKey, value := range m[pop.s] {
				if sKey == v[1] {
					res = append(res, pop.v*value)
					break
				}
				q = append(q, Item{sKey, pop.v * value})
			}
		}
		if len(res) < k+1 {
			res = append(res, -1.0)
		}
	}
	return res
}
