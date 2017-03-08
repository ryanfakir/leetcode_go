package main

func findWords(words []string) []string {
	Agroup := []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	Bgroup := []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'}
	Cgroup := []rune{'z', 'x', 'c', 'v', 'b', 'n', 'm'}
	var res []string
	for _, v := range words {
		var a, b, c int
		for _, point := range v {
			if point < 'a' {
				point += 32
			}
			for _, el := range Agroup {
				if el == point {
					a = 1
				}
			}
			for _, el := range Bgroup {
				if el == point {
					b = 4
				}
			}
			for _, el := range Cgroup {
				if el == point {
					c = 8
				}
			}
		}
		test := a | b | c
		if test&(test-1) == 0 {
			res = append(res, v)
		}
	}
	return res
}
