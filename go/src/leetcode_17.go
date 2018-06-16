package main

import "strconv"

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
func letterCombinations(digits string) []string {
	dict := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var res []string
	for _, val := range digits {
		num, _ := strconv.Atoi(string(val))
		var temp []string
		for _, digit := range dict[num-2] {
			if len(res) == 0 {
				temp = append(temp, string(digit))
			}
			for _, ins := range res {
				temp = append(temp, ins+string(digit))
			}
		}
		res = temp
	}
	return res
}
func letterCombinations(digits string) []string {
    if len(digits) ==0 {return []string{}}
    dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno","pqrs", "tuv", "wxyz"}
    res := []string{""}
    for _, v := range digits {
        tmp := dict[v-'0']
        var newRes []string
        for _, t := range tmp {
            for _, r := range res {
                newRes = append(newRes, r + string(t))
            }
        }
        res = newRes
    }
    return res
}
