package problems

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, v := range s {
		switch v {
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		default:
			stack = append(stack, byte(v))
		}
	}
	println(len(stack))
	return len(stack) == 0

}
