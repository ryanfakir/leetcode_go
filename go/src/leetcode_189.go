package src

func rotate(nums []int, k int)  {
    if len(nums) == 0 {
        return
    }
    if k > len(nums) {
        k = k%len(nums)
    }
    reverse(nums[:len(nums)-k])
    reverse(nums[len(nums)-k:])
    reverse(nums)
}

func reverse(arr []int) {
    for i,j := 0,len(arr)-1; i< j; i,j = i+1, j-1 {
        arr[i], arr[j] = arr[j],arr[i]
    }
}
