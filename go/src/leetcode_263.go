package src

func isUgly(num int) (res bool) {
	if num < 0 {
		return
	}
	temp := num

	for temp > 0 {
		if temp%2 == 0 {
			temp /= 2
		} else if temp%3 == 0 {
			temp /= 3
		} else if temp%5 == 0 {
			temp /= 5
		} else if temp == 1 {
			res = true
			break
		} else {
			res = false
			break
		}
	}
	return
}
