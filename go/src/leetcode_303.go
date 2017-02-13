package main

type NumArray struct {
	InnerArr []int
}

func Constructor(nums []int) NumArray {
	var instance NumArray
	var arr []int
	var temp int
	for _, v := range nums {
		temp += v
		arr = append(arr, temp)
	}
	instance.InnerArr = arr
	return instance
}

func (this *NumArray) SumRange(i int, j int) int {
	if i-1 < 0 {
		return this.InnerArr[j]
	}
	return this.InnerArr[j] - this.InnerArr[i-1]
}
