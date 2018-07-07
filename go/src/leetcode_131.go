package main

func partition(s string) [][]string {
	return helper(0, s, nil, nil)
}

func helper(level int, s string, temp []string, res [][]string) [][]string {
	if level == len(s) {
		el := make([]string, len(temp))
		copy(el, temp)
		res = append(res, temp)
		return res
	}
	for i := level + 1; i <= len(s); i++ {
		if valid(s[level:i]) {
			//fmt.Println(level)
			//fmt.Println(s[level: i])
			temp := append(temp, s[level:i])
			res = helper(i, s, temp, res)
			temp = temp[:len(temp)-1]
		}
	}
	return res
}

func valid(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	return helper(0, s, nil, nil)
}

func helper(level int, s string, temp []string, res [][]string) [][]string {
	if level == len(s) {
		el := make([]string, len(temp))
		copy(el, temp)
		res = append(res, temp)
		return res
	}
	org := temp
	for i := level + 1; i <= len(s); i++ {
		if valid(s[level:i]) {
			//fmt.Println(level)
			//fmt.Println(s[level: i])
			org = temp
			temp = append(temp, s[level:i])
			res = helper(i, s, temp, res)
			temp = org
		}
	}
	return res
}

func valid(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
    return dfs(s, nil, nil, 0)
}


func dfs(s string, tmp []string, res [][]string, level int ) [][]string {
    if level == len(s) {
        in := make([]string, len(tmp))
        copy(in, tmp)
        res = append(res, in)
        return res
    }
    for i := level+1; i <= len(s); i++ {
        if helper(s, level, i-1) {
            tmp = append(tmp, s[level:i])
            res = dfs(s, tmp, res, i)
            tmp = tmp[:len(tmp)-1]
        }
    }
    return res
}

func helper(s string, i, j int) bool {
    if i > j { return false}
    for i <= j {
        if s[i] != s[j] {return false}
        i++
        j--
    }
    return true
}
