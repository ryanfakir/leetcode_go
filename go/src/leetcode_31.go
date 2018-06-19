package main

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1, len(nums)-1)
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func nextPermutation(nums []int)  {
    i := len(nums)- 2
    for i >=0 && nums[i] >= nums[i+1] {
        i--
    }
    j := i+1
    for i >= 0 && j < len(nums) && nums[j] > nums[i] {
        j++
    }
    j--
    if i >=0 && j >= 0 {
         tmp := nums[j]
        nums[j] = nums[i]
        nums[i] = tmp
    }
    reverse(nums, i+1, len(nums)-1)
}


func reverse(nums []int, i, j int) {
    for i < j {
        tmp := nums[i]
        nums[i] = nums[j]
        nums[j] = tmp
        i++
        j--
    }
}
