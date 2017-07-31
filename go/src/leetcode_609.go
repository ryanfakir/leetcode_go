package main

import "strings"

func findDuplicate(paths []string) [][]string {
	dict := make(map[string][]string)
	for _, v := range paths {
		temp := strings.Split(v, " ")
		var pre string
		for i, path := range temp {
			if i == 0 {
				pre = path
				continue
			}
			index := strings.Index(path, "(")
			post := pre + "/" + path[:index]
			content := path[index+1 : len(path)-1]
			dict[content] = append(dict[content], post)
		}
	}
	var res [][]string
	for _, v := range dict {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}
