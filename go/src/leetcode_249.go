func groupStrings(strings []string) [][]string {
    dict := make(map[string][]string)
    for _, str := range strings {
        var t string
        for _ , s := range str {
            t += strconv.Itoa(int(s + 26 - rune(str[0])) % 26) + ","  
        }
        dict[t] = append(dict[t], str)
    }
    var res [][]string
    for _, v := range dict {
        res = append(res, v)
    }
    return res
}
