package src

func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers)-1
    for left < right {
        if (numbers[left] + numbers[right] > target) {
            right--
        } else if (numbers[left] + numbers[right] < target) {
            left++
        } else {
            res := make([]int, 0)
            res = append(res, left+1)
            res = append(res, right+1)
            return res
        }

    }
    return make([]int, 2)
}
