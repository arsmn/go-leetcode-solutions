package leetcodesolutions

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	var (
		stack    = make([]rune, 0)
		brackets = map[rune]rune{
			'(': ')',
			'{': '}',
			'[': ']',
		}
	)

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}

			if brackets[stack[len(stack)-1]] != r {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
