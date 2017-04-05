package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(Item(people))
	res := make([][]int, len(people))
	fmt.Println(people)
	for _, v := range people {
		res[v[1]] = v
	}
	return res

}

type Item [][]int

func (el Item) Less(i, j int) bool {
	return el[i][0] > el[j][0] || (el[i][0] == el[j][0] && el[i][1] < el[j][1])
}

func (el Item) Len() int { return len(el) }

func (el Item) Swap(i, j int) { el[i], el[j] = el[j], el[i] }
