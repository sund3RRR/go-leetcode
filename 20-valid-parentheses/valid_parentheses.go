// 20. Valid Parentheses
//
// Difficulty: Easy
// Link:	   https://leetcode.com/problems/valid-parentheses
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//   - Open brackets must be closed by the same type of brackets.
//   - Open brackets must be closed in the correct order.
//   - Every close bracket has a corresponding open bracket of the same type.
package valid_parentheses

func IsValid(s string) bool {
	return isValid(s)
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, r := range s {
		if isOpen(r) {
			stack = append(stack, revert(r))
		} else if len(stack) > 0 && r == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func isOpen(r rune) bool {
	switch r {
	case '(', '{', '[':
		return true
	}
	return false
}

func revert(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		return '('
	}
}
