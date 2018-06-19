func searchRange(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right)/2
        if target == nums[mid] {
            fmt.Println(mid)
            ll, lr := 0, mid
            // lm = 2
            for ll <= lr {
                lm := (ll + lr )/ 2
                if nums[lm] == target {
                    lr = lm -1
                } else {
                    ll = lm +1
                }
            }
            lb := ll
            rl, rr := mid, len(nums)-1
            for rl <= rr {
                rm := (rl + rr) /2
                if nums[rm] == target {
                    rl = rm +1
                } else {
                    rr = rm -1
                }
            }
            rb := rr
            return []int{lb, rb}
        } else if target < nums[mid] {
            right = mid -1
        } else {
            left = mid +1
        }
    }
    return []int{-1, -1}
}
