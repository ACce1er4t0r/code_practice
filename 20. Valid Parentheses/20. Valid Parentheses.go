package isValid

func isValid(s string) bool {
	dic := map[rune]rune{
		40:  41,
		91:  93,
		123: 125,
	}
	stack := []rune{}
	for _, v := range s {
		if v == 40 || v == 91 || v == 123 {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			l := len(stack)
			top := stack[l-1]
			stack = stack[:l-1]
			if dic[top] != v {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
