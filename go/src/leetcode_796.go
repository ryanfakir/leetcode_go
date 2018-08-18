func rotateString(A string, B string) bool {
    if len(A) == len(B) {
        return strings.Contains(A+A, B)
    }
    return false
}
