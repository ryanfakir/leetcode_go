package src

func hammingDistance(x int, y int) int {
	res := x ^ y
	var total int
	for res != 0 {
		res &= res - 1
		total++
	}
	return total
}
