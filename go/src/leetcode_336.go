func palindromePairs(words []string) [][]int {
    dict := make(map[string]int)
    for index, v := range words {
        arr := strings.Split(v, "")
        for i, j := 0, len(arr)-1; i< j; i, j = i+1, j-1 {
            arr[i], arr[j] = arr[j], arr[i]
        }
        v = strings.Join(arr, "")
        dict[v] = index + 1
    }
    var res [][]int
    for i, word := range words {
        for j:= 0; j <=len(word); j++ {
            left, right := word[0:j], word[j:]
            if valid(left) && dict[right] != 0 && dict[right] != i +1 &&  len(word) > len(word) - j{
                res = append(res, []int{dict[right]-1, i})
            }
            if valid(right) && dict[left] != 0 && dict[left] != i+1 {
                res = append(res, []int{i, dict[left]-1})
            }
        }
    }
    return res
}

func valid(s string) bool {
    for i, j := 0, len(s)-1; i < j ; i, j = i+1, j-1 {
        if s[i] != s[j] {return false}
    }
    return true
}
