package main

import "fmt"

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
		fmt.Println(q)
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
						fmt.Println(m)
					}
				}
			}
		}
	}
	return 0
}
