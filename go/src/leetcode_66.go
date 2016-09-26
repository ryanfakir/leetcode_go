package src

func plusOne(digits []int) []int {
    length := len(digits)
    carry := 1
    for i:= length-1; i> -1; {
        if carry != 0 {
            digits[i] += 1
            carry = 0
        }
        if digits[i] > 9 {
            digits[i] = 0
            carry = 1
        }
        i--
    }
    if carry == 1 {
        res := make([]int, 1)
        res[0] = 1
        digits = append(res, digits...)
    }
    return digits
}
