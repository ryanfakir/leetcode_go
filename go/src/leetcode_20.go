package src

func isValid(s string) bool {
	dict := make(map[rune]rune, 0)
	dict['}'] = '{'
	dict[')'] = '('
	dict[']'] = '['
	stack := make([]rune, 0)
	for _, e := range s {
		if len(stack) == 0 {
			stack = append(stack, e)
		} else if dict[e] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, e)
		}
	}
	return len(stack) == 0
}
func isValid(s string) bool {
	dict := make(map[rune]rune, 0)
	dict['}'] = '{'
	dict[')'] = '('
	dict[']'] = '['
	var stack []rune
	for _, e := range s {
		if len(stack) == 0 {
			stack = append(stack, e)
		} else if dict[e] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, e)
		}
	}
	return len(stack) == 0
}

func isValid(s string) bool {
    var stack []rune
    for _ , v := range s {
        if v == '(' || v == '[' || v == '{' {
            if v == '(' {stack = append(stack, ')')}
            if v == '{' {stack = append(stack, '}')}
            if v == '[' {stack = append(stack, ']')}
            continue
        }
        if v == ')' || v == ']' || v == '}' {
            if len(stack) == 0 {return false}
            pop := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if pop != v {return false}
        }
    }
    return len(stack) == 0
}
