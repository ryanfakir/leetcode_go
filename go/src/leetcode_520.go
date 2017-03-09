package main

func detectCapitalUse(word string) bool {
	var indicator int
	var cap bool
	for i := 0; i < len(word); i++ {
		if i == 0 {
			if word[i] > 'Z' {
				indicator = 1
			}
			continue
		}
		if indicator == 1 {
			if word[i] <= 'Z' {
				return false
			}
		} else {
			if i == 1 {
				if word[i] <= 'Z' {
					cap = true
				}
				continue
			}
			if cap {
				if word[i] > 'Z' {
					return false
				}
			} else {
				if word[i] <= 'Z' {
					return false
				}
			}
		}
	}
	return true
}
