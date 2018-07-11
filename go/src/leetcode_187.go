func findRepeatedDnaSequences(s string) []string {
    m := make(map[string]int)
    var res []string
    for i :=0 ; i <= len(s)-10; i++ {
        fmt.Println(i)
        if m[s[i:i+10]] == 1 {
            res = append(res, s[i:i+10])
        }
        m[s[i:i+10]]++
    }
    return res
}
