package src
func isIsomorphic(s string, t string) bool {
    sMap := make(map[byte]byte, 0)
    tMap := make(map[byte]byte, 0)
    for i:=0; i< len(s); i++ {
        if e, ok := sMap[s[i]]; ok {
            if e != t[i] {
                return false
            }
        } else {
            sMap[s[i]] = t[i]
        }

        if e, ok := tMap[t[i]]; ok {
            if e != s[i] {
                return false
            }
        } else {
            tMap[t[i]] = s[i]
        }
    }
    return true
}
