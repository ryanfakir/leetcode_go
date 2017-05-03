package main

func checkSubarraySum(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				return true
			}
			if k != 0 && sum%k == 0 {
				return true
			}
		}
	}
	return false
}
