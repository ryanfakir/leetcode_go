type Solution struct {
    sum []int
    total int
}


func Constructor(w []int) Solution {
    var total int
    var sum []int
    for _, v := range w {
        total += v
        sum = append(sum, total)
    }
    return Solution{sum: sum, total: total}
}


func (this *Solution) PickIndex() int {
    ran := rand.New(rand.NewSource(time.Now().UnixNano()))
    num := ran.Intn(this.total)
    l, r := 0, len(this.sum)-1
    for l < r {
        mid := (l + r) /2
        if this.sum[mid] > num {
            r = mid
        } else {
            l = mid +1
        }
    }
    return l
}

