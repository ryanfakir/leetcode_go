package main

import (
	"math"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Sort(PointArr(points))
	var arr [][]int
	for _, point := range points {
		var tempArr [][]int
		for _, temp := range arr {
			if temp[0] > point[1] || temp[1] < point[0] {
				tempArr = append(tempArr, temp)
				continue
			}
			point = []int{int(math.Max(float64(point[0]), float64(temp[0]))), int(math.Min(float64(point[1]), float64(temp[1])))}
		}
		tempArr = append(tempArr, point)
		arr = tempArr
	}
	return len(arr)
}

type PointArr [][]int

func (p PointArr) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p PointArr) Len() int {
	return len(p)
}

func (p PointArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(PointArr(points))
	var res int = 1
	high := points[0][1]
	for i, point := range points {
		if i == 0 {
			continue
		}
		if point[0] <= high {
			high = int(math.Min(float64(high), float64(point[1])))
		} else {
			res++
			high = point[1]
		}
	}
	return res
}

type PointArr [][]int

func (p PointArr) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p PointArr) Len() int {
	return len(p)
}

func (p PointArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
