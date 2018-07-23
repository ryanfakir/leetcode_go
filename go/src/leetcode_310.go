package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	container := make(map[int]map[int]bool)
	for _, v := range edges {
		if len(container[v[0]]) == 0 {
			container[v[0]] = make(map[int]bool)
		}
		if len(container[v[1]]) == 0 {
			container[v[1]] = make(map[int]bool)
		}
		container[v[0]][v[1]] = true
		container[v[1]][v[0]] = true
	}
	var q []int
	for k, v := range container {
		var i int
		for _ = range v {
			i++
		}
		if i == 1 {
			q = append(q, k)
		}
	}
	//fmt.Println(container)
	for n > 2 {
		length := len(q)
		n -= length
		for i := 0; i < length; i++ {
			pop := q[0]
			q = q[1:]
			for k, v := range container {
				if v[pop] {
					v[pop] = false
					var i int
					for _, value := range v {

						if value == true {
							i++
						}
					}
					if i == 1 {
						q = append(q, k)
					}
				}
			}
			container[pop] = nil
		}
	}
	//fmt.Println(container)
	var res []int
	for _, v := range q {
		res = append(res, v)
	}
	return res
}

func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {return []int{0}}
    dict := make(map[int][]int)
    for _, v := range edges {
        dict[v[0]] = append(dict[v[0]], v[1])
        dict[v[1]] = append(dict[v[1]], v[0])
    }
    var q []int
    for k, v := range dict {
        if len(v) == 1 {
            q = append(q, k)
        }
    }
    for len(q) > 0 {
        size := len(q)
        if n <= 2 {break}
        for i:= 0; i < size; i++ {
            pop := q[0]
            q = q[1:]
            n--
            p := dict[pop][0]
            for i, v := range dict[p] {
                if v == pop {
                    dict[p] = append(dict[p][:i], dict[p][i+1:]...)
                    if len(dict[p]) == 1 {
                        q =append(q, p)
                    }
                   
                }
            }
            delete(dict, pop)
        }
    }
    var res []int
    for _, v := range q {
        res = append(res, v)
    }
    return res
}
