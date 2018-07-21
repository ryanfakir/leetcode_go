type MedianFinder struct {
    small *PQue
    big *PQue
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    return MedianFinder{small:&PQue{}, big: &PQue{}}
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.small, float64(-num))
    heap.Push(this.big, -heap.Pop(this.small).(float64))
    if this.small.Len() < this.big.Len() {
        heap.Push(this.small, -heap.Pop(this.big).(float64))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.big.Len() == this.small.Len() {
        return ((*(this.big))[0] - (*(this.small))[0]) /2
    } else {
        return -(*(this.small))[0]     
    }
}


type PQue []float64

func (pq PQue) Less(i, j int) bool {
    return pq[i] < pq[j]
}

func (pq PQue) Len() int {
    return len(pq)
}

func (pq PQue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQue) Push(item interface{}) {
    *pq = append(*pq, item.(float64))
}

func (pq *PQue) Pop() interface{} {
    tmp := *pq
    pop := tmp[len(tmp)-1]
    *pq = tmp[:len(tmp)-1]
    return pop
}
