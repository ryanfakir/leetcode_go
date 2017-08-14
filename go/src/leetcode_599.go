package main

func findRestaurant(list1 []string, list2 []string) []string {
	dict := make(map[string]int)
	for i, v := range list1 {
		dict[v] = i + 1
	}
	var res []string
	var sum = 1 << 11
	for i, v := range list2 {
		if dict[v] != 0 {
			if i+dict[v] > sum {
				continue
			}
			if i+dict[v] < sum {
				res = []string{}
				sum = i + dict[v]
			}
			res = append(res, v)
		}
	}
	return res
}
