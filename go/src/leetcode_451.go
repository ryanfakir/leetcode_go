func frequencySort(s string) string {
    dict := make(map[rune]int)
    for _ , v := range s {
        dict[v]++
    }
    pq := new(PQ)
    heap.Init(pq)
    for k, v := range dict {
        heap.Push(pq, Entry{key: k, val: v})
    }
    var res string
    for pq.Len() > 0 {
        pop := heap.Pop(pq).(Entry)
        for pop.val > 0 {
            res += string(pop.key)
            pop.val--
        }
    }
    return res
}

type PQ []Entry

type Entry struct {
    key rune
    val int
}

func (pq *PQ) Push(el interface{}) {
    *pq = append(*pq, el.(Entry))
}

func (pq *PQ) Pop() interface{} {
    arr := *pq
    pop := arr[len(arr)-1]
    arr = arr[:len(arr)-1]
    *pq = arr
    return pop 
}

func (pq PQ) Len() int {return len(pq)}
func (pq PQ) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]}
func (pq PQ) Less(i, j int) bool {return pq[i].val > pq[j].val}
