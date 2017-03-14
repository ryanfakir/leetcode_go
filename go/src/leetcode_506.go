package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	length := len(nums)
	pq := make(PriorityQueue, length)
	heap.Init(&pq)
	for i, v := range nums {
		heap.Push(&pq, &Item{i, v})
	}
	fmt.Println("temp")
	res := make([]string, length)
	cnt := 1
	for i := 0; i < length; i++ {
		item := heap.Pop(&pq).(*Item)

		if i == 0 {
			res[item.value] = "Gold Medal"
		} else if i == 1 {
			res[item.value] = "Silver Medal"
		} else if i == 2 {
			res[item.value] = "Bronze Medal"
		} else {
			res[item.value] = strconv.Itoa(cnt)
		}
		cnt++
	}
	return res
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
