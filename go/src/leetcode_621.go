package main

func leastInterval(tasks []byte, n int) int {
	var max, window int
	dict := make([]int, 26)
	for _, v := range tasks {
		dict[rune(v)-'A']++
		if dict[rune(v)-'A'] > max {
			max = dict[rune(v)-'A']
			window = 1
		} else if dict[rune(v)-'A'] == max {
			window++
		}
	}
	var group = max - 1
	var empty = n - (window - 1)
	var tEmpty = group * empty
	var left = len(tasks) - max*window
	var idle int
	if tEmpty-left > idle {
		idle = tEmpty - left
	}
	return len(tasks) + idle
}
