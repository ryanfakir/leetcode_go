package src

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start int, end int) *TreeNode {
    mid := (end-start)/2 + start
    root := new(TreeNode)
    root.Val = nums[mid]
    if start <= mid-1 {
        root.Left = helper(nums, start, mid-1)
    }
    if mid+1 <= end {
        root.Right = helper(nums, mid+1, end)
    }
    return root
}
