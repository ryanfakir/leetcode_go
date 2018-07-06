package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var q []string
	m := make(map[string]int)
	q = append(q, beginWord)
	var valid bool
	for _, v := range wordList {
		if endWord == v {
			valid = true
		}
	}
	if !valid {
		return 0
	}
	m[beginWord] = 1
	visited := make(map[string]bool)
	for len(q) > 0 {
		pop := q[0]
		q = q[1:]
		if _, ok := visited[pop]; ok {
			continue
		}
		visited[pop] = true
		for i := 0; i < len(pop); i++ {
			temp := []rune(pop)
			for j := 'a'; j <= 'z'; j++ {
				temp[i] = j
				//fmt.Println(string(temp))
				if string(temp) == endWord {
					return m[pop] + 1
				}
				for _, v := range wordList {
					if v == string(temp) && !visited[v] {
						q = append(q, v)
						m[v] = m[pop] + 1
					}
				}
			}
		}
	}
	return 0
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	var q []string
	q = append(q, beginWord)
	visited := make(map[string]bool)
	level := make(map[string]bool)
	step := 1
	for len(q) > 0 {
		l := len(q)
		step++
		for i := 0; i < l; i++ {
			pop := q[0]
			q = q[1:]
			if visited[pop] {
				continue
			}
			visited[pop] = true
			for j := 0; j < len(pop); j++ {
				for k := 'a'; k <= 'z'; k++ {
					temp := pop[0:j] + string(k) + pop[j+1:]
					if visited[temp] {
						continue
					}
					if dict[temp] && !level[temp] {
						q = append(q, temp)
						level[temp] = true
						if temp == endWord {
							return step
						}
					}
				}
			}
		}
	}
	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    dict := make(map[string]bool)
    for _, v := range wordList {
        dict[v] = true
    }
    var q []string = []string{beginWord}
    var level = 1
    gMap := make(map[string]bool)
    for len(q) > 0 {
        size := len(q)
        levelMap := make(map[string]bool)
        for i:= 0; i< size; i++ {
            pop := q[0]
            q = q[1:]
            gMap[pop] = true
            if pop ==endWord {return level}
            for j :=0; j < len(pop); j++ {
                tmp := pop
                arr := []byte(tmp)
                for k := 'a'; k <= 'z'; k++ {
                    arr[j] = byte(k)
                    if !gMap[string(arr)] && !levelMap[string(arr)] && dict[string(arr)] {
                        fmt.Println(string(arr))
                        levelMap[string(arr)] = true
                        q = append(q, string(arr))
                    }
                }
            }
        }
        level++
    }
    return 0
}
