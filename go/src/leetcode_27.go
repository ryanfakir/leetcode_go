package src

func removeElement(nums []int, val int) int {
    b := nums[:0]
    for _, e := range nums {
        if e != val {
            b = append(b, e)
        }
    }
    return len(b)
}
