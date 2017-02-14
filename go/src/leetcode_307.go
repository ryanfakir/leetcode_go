package main

import "math"

type NumArray struct {
	InnerArr []int
	tail     int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	size := int(math.Pow(2, math.Ceil(math.Log2(float64(len(nums))))))
	instance := &NumArray{make([]int, size*2-1), len(nums)}
	instance.cTree(0, len(nums)-1, 0, nums)
	return *instance
}

func (this *NumArray) cTree(start, end, pos int, nums []int) {
	if start == end {
		this.InnerArr[pos] = nums[start]
		return
	}
	mid := start + (end-start)/2
	this.cTree(start, mid, 2*pos+1, nums)
	this.cTree(mid+1, end, 2*pos+2, nums)
	this.InnerArr[pos] = this.InnerArr[2*pos+1] + this.InnerArr[2*pos+2]
}

func (this *NumArray) Update(i int, val int) {
	this.helper(i, val, 0, this.tail-1, 0)
}

func (this *NumArray) helper(i, val, start, end, pos int) {
	if start == end {
		if start == i {
			this.InnerArr[pos] = val
		}
		return
	}
	mid := start + (end-start)/2

	if mid >= i {
		this.helper(i, val, start, mid, 2*pos+1)
		this.InnerArr[pos] = this.InnerArr[2*pos+1] + this.InnerArr[2*pos+2]
	} else if mid < i {
		this.helper(i, val, mid+1, end, 2*pos+2)
		this.InnerArr[pos] = this.InnerArr[2*pos+1] + this.InnerArr[2*pos+2]
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.rangeHelper(i, j, 0, this.tail-1, 0)
}

func (this *NumArray) rangeHelper(i, j, start, end, pos int) int {
	if i <= start && j >= end {
		return this.InnerArr[pos]
	}
	if i > end || j < start {
		return 0
	}
	mid := start + (end-start)/2
	if mid < i {
		return this.rangeHelper(i, j, mid+1, end, 2*pos+2)
	} else if mid > j {
		return this.rangeHelper(i, j, start, mid, 2*pos+1)
	} else {
		return this.rangeHelper(i, j, mid+1, end, 2*pos+2) + this.rangeHelper(i, j, start, mid, 2*pos+1)
	}
}
