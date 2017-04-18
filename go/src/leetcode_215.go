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
