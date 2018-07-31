package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	start, end := 0, len(nums)-1
	for {
		index := helper(nums, start, end)
		if index > k {
			end = index - 1
		} else if index < k {
			start = index + 1
		} else {
			break
		}
	}
	return nums[k]
}

func helper(nums []int, start, end int) int {
	pvIndex := end
	pv := nums[end]
	end--
	for start <= end {
		for start <= end && nums[start] < pv {
			start++
		}
		for start <= end && nums[end] > pv {
			end--
		}
		if start <= end {
			temp := nums[start]
			nums[start] = nums[end]
			nums[end] = temp
			start++
			end--
		}
	}
	temp := nums[pvIndex]
	nums[pvIndex] = nums[start]
	nums[start] = temp
	return start
}
func findKthLargest(nums []int, k int) int {
	// h := make(IntElement, len(nums))
	var h IntElement
	heap.Init(&h)
	for _, v := range nums {
		// h[i] = v
		heap.Push(&h, v)
	}
	// heap.Init(&h)
	for len(nums)-k > 0 {
		k++
	}
	return heap.Pop(&h).(int)
}

type IntElement []int

func (el IntElement) Swap(i, j int)      { el[i], el[j] = el[j], el[i] }
func (el IntElement) Less(i, j int) bool { return el[i] < el[j] }
func (el IntElement) Len() int           { return len(el) }

func (p *IntElement) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *IntElement) Pop() interface{} {
	temp := *p
	el := temp[len(temp)-1]
	*p = temp[:len(temp)-1]
	return el
}

func findKthLargest(nums []int, k int) int {
    var pq ElementArr
    heap.Init(&pq)
    for _, v := range nums {
        heap.Push(&pq, v)
    }
    for  k > 1 {
        heap.Pop(&pq)
        k--
    }
    return heap.Pop(&pq).(int)
}

type ElementArr []int

func (e ElementArr) Len() int {return len(e)}
func (e ElementArr) Swap(i, j int) {e[i], e[j] = e[j], e[i]}
func (e ElementArr) Less(i, j int) bool {return e[i] > e[j]}
func (e *ElementArr) Push(input interface{}) {
    *e = append(*e, input.(int))
}

func (e *ElementArr) Pop() interface{} {
    temp := *e
    res := temp[len(temp) -1]
    *e = temp[:len(temp)-1]
    return res
}



func findKthLargest(nums []int, k int) int {
    pq := &PQ{}
    heap.Init(pq)
    for _ , v := range nums {
        heap.Push(pq, v)
        if pq.Len() > k {
            heap.Pop(pq)
        }
    }
    return heap.Pop(pq).(int)
}


type PQ []int

func (p *PQ) Push(x interface{}) {
    *p = append(*p, x.(int))
}

func (p *PQ) Pop() interface{} {
    org := *p
    pop := org[len(org)-1]
    *p = org[:len(org)-1]
    
    return pop
}

func (p PQ) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p PQ) Len() int {return len(p)}
func (p PQ) Less(i, j int) bool {return p[i] < p[j]}
