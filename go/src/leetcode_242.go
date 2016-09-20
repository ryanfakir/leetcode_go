package src

func isAnagram(s string, t string) bool {
    if (len(s) != len(t)) {
        return false
    }
    dict := make([]int, 26)
    res := true
    for _, e := range s {
        dict[e-'a']++
    }
    for _, e := range t {
        if dict[e-'a'] == 0 {
            res = false
            break
        }
        dict[e-'a']--
    }
    return res
}
