func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 {
			sum[i] = v
			continue
		}
		sum[i] = sum[i-1] + v
	}
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			var prev int
			if i-1 < 0 {
				prev = 0
			} else {
				prev = sum[i-1]
			}
			if sum[j]-prev == k {
				res++
			}
		}
	}
	return res
}
