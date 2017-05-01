package main

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	var res []int
	cnt := make([]int, 26)
	for _, v := range p {
		cnt[v-'a']++
	}
	for i := 0; i < len(s)-len(p)+1; i++ {
		temp := make([]int, 26)
		copy(temp, cnt)
		good := true
		for j := i; j < i+len(p); j++ {
			temp[s[j]-'a']--
			if temp[s[j]-'a'] < 0 {
				good = false
				break
			}
		}
		if good {
			res = append(res, i)
		}
	}
}
