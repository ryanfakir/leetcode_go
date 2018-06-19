func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right) / 2 
        if nums[mid] < nums[left] {
            if target > nums[right] || target < nums[mid]   {
                right = mid -1
            } else {
                left = mid + 1
            }
        }else if nums[mid] > nums[right] {
            if target < nums[left] || target > nums[mid]  {
                left = mid + 1
            } else {
                right = mid -1
            }
        } else {
            if target > nums[mid] {
                left = mid + 1
            } else {
                right = mid -1
            }
        } 
        if nums[mid] == target {
            return mid
        }
        
    }
    return -1
}
