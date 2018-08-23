func smallestDistancePair(nums []int, k int) int {
    sort.Ints(nums)
    left, right := 0, nums[len(nums)-1] - nums[0]
    for left < right{
        mid := (left + right)/2
        var cnt int
        for i:= 0; i < len(nums); i++ {
            j:= i +1
            for j < len(nums)&& nums[j]- nums[i] < mid {j++}
            cnt += j - i - 1
        }
        fmt.Println(cnt, " :cnt ", mid)
        fmt.Println(left, right)
        if cnt < k {
            left = mid +1 
        } else {
            right = mid 
        }
    }
    return right
}
