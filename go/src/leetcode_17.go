package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var dict = []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res = []string{""}
	for _, v := range digits {
		var temp []string
		for _, element := range dict[v-'0'] {
			for _, val := range res {
				temp = append(temp, val+string(element))
			}
		}
		res = temp
	}
	return res
}
