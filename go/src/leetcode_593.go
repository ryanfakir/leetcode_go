package main

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var dict = [][]int{p1, p2, p3, p4}
	var set = make(map[int]bool)
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			set[(dict[i][0]-dict[j][0])*(dict[i][0]-dict[j][0])+(dict[i][1]-dict[j][1])*(dict[i][1]-dict[j][1])] = true
		}
	}
	for k, _ := range set {
		if k == 0 {
			return false
		}
	}
	return len(set) == 2
}
