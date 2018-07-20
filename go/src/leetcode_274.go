func hIndex(citations []int) int {
    sort.Ints(citations)
    for i := len(citations)-1; i >= 0; i-- {
        if len(citations) -1 - i >= citations[i] {
            return len(citations) -1 - i
        } 
    }
    return len(citations)
}
