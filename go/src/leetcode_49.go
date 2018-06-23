package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, v := range strs {
		org := v
		temp := strings.Split(v, "")
		sort.Strings(temp)
		v = strings.Join(temp, "")
		m[v] = append(m[v], org)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	sort.Strings(strs)
	m := make(map[string][]string)
	for _, v := range strs {
		org := v
		temp := strings.Split(v, "")
		sort.Strings(temp)
		v = strings.Join(temp, "")
		m[v] = append(m[v], org)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
    dict := make(map[string][]string)
    for _, v := range strs {
        arr :=strings.Split(v, "")
        sort.Strings(arr)
        tmp := strings.Join(arr, "")
        dict[tmp] = append(dict[tmp], v)
    }
    var res [][]string
    for _, v := range dict {
        res = append(res, v)
    }
    return res
}
