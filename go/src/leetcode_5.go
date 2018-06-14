package main

func longestPalindrome(s string) string {
	var res string
	for i, _ := range s {
		left, right := i, i
		odd := helper(left, right, s)
		even := helper(left, right+1, s)
		if len(odd) >= len(res) {
			res = odd
		}
		if len(even) >= len(res) {
			res = even
		}
	}
	return res
}

func helper(left int, right int, s string) (res string) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			if len(s[left:right+1]) >= len(res) {
				res = s[left : right+1]
			}
			left--
			right++
			continue
		}
	}
	return
}

func longestPalindrome(s string) string {
    var res string
    for i := range s {
        single := isPalindrome(s, i, i)
        double := isPalindrome(s, i, i+1)
        if len(single) > len(double) {
            if len(single) > len(res) {
                res = single
            }
        } else {
            if len(double) > len(res) {
                res = double
            }
        }
    }
    return res
}

func isPalindrome(s string, left, right int) string {
    var res string
    for left >= 0 && right < len(s) {
        if s[left] != s[right] {
            break
        }
        res = s[left: right+1]
        left--
        right++
    }
    return res
}
