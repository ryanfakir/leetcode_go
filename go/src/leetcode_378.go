package src

import "container/heap"
func kthSmallest(matrix [][]int, k int) int {
    len := len(matrix)
    pq := new(IntHeap)
    heap.Init(pq)
    for _,row := range matrix {
        for _,el := range row {
            heap.Push(pq, el)
            if pq.Len() > len*len - k +1 {
                heap.Pop(pq)
            }
        }
    }
    return heap.Pop(pq).(int)
}

type IntHeap []int

func (pq IntHeap) Len() int {return len(pq)}
func (pq IntHeap) Less(i, j int) bool {return pq[i]<pq[j]}
func (pq IntHeap) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]}

func (pq *IntHeap) Push(x interface{}) {
    *pq = append(*pq, x.(int))
}

func (pq *IntHeap) Pop() interface{} {
    n := len(*pq)
    x := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return x
}

