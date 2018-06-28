func search(nums []int, target int) bool {
    if len(nums) == 0 {return false}
    start, end := 0, len(nums)-1
    for start <= end {
        fmt.Println(start, " ", end)
        mid := (start + end)/2
         if nums[mid] == target {
             return true
         } else if nums[mid] > nums[end] {
            if target >= nums[mid] || target <= nums[end] {
                start = mid +1
            } else {
                end = mid -1
            }
        } else if nums[mid] < nums[start] {
            if target <= nums[mid] || target >= nums[start] {
                end = mid -1
            } else {
                start = mid+ 1
            }
        } else {
             if nums[start] == target || nums[end] == target {return true}
            start++
            end--
        }
    }
    return false
}
