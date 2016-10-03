package src

import "strings"
func lengthOfLastWord(s string) int {
    res := 0
    s  = strings.TrimSpace(s)
    for _, e := range s {
        if e == ' ' {
            res = 0
        } else {
            res++
        }
    }
    return res
}
