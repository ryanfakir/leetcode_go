func generatePossibleNextMoves(s string) []string {
    if len(s) <= 1 {return nil}
    var res []string
    for i:= 1; i < len(s); i++ {
        if s[i] == '+' && s[i-1] == '+' {
            res = append(res, s[:i-1] + "--" + s[i+1:])
        }
    }
    return res
}
