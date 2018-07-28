package src

import "container/heap"
func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int, 0)
    for _, el := range nums {
        if val, ok := m[el]; ok {
            m[el] = val +1
        } else {
            m[el] = 1
        }
    }

    pq := make(IntHeap, 0)
    heap.Init(&pq)
    for key, v := range m {
        entry := &Entry{key, v}
        heap.Push(&pq, entry)
        if pq.Len() > k {
            heap.Pop(&pq)
        }
    }
    res := make([]int, 0)

    for pq.Len() > 0 {
        el := heap.Pop(&pq).(*Entry)
        res = append(res, el.key)
    }
    return res
}
type Entry struct {
    key int
    value int
}
type IntHeap []*Entry

func (q *IntHeap) Push(x interface{}) {
    *q = append(*q, x.(*Entry))
}

func (q *IntHeap) Pop() interface{} {
    old := *q
    n:=len(old)
    x:=old[n-1]
    *q = old[: n-1]
    return x
}

func (q IntHeap) Len() int { return len(q)}
func (q IntHeap) Less(i, j int) bool { return q[i].value < q[j].value}
func (q IntHeap) Swap(i, j int) {q[i], q[j] = q[j], q[i]}

func topKFrequent(nums []int, k int) []int {
    dict := make(map[int]int)
    for _ , v := range nums {
        dict[v]++
    }
    var pq PQ
    heap.Init(&pq)
    for key, v := range dict {
        if len(pq) >= k {
            pop := heap.Pop(&pq).(Entry)
            if pop.value > v {
                 heap.Push(&pq, pop)
                continue
            }
        }
        heap.Push(&pq, Entry{key, v})
    }
    var res []int
    for pq.Len() > 0{
        res = append(res, heap.Pop(&pq).(Entry).key)
    }
    return res
}
