package problems

// Given a string containing just the characters '(' and ')',
// find the length of the longest valid (well-formed) parentheses substring.
//
// Example 1:
//
//	Input: "(()"
//	Output: 2
//	Explanation: The longest valid parentheses substring is "()"
//
// Example 2:
//
//	Input: ")()())"
//	Output: 4
//	Explanation: The longest valid parentheses substring is "()()"
func LongestValidParentheses(s string) int {
	size := len(s)
	stack := NewStack(size)
	start := 0
	res := 0
	for i, char := range s {
		if char == '(' {
			stack.Push(uint8(i))
		} else {
			if stack.Size() == 0 {
				start = i + 1
			} else if stack.Size() == 1 {
				stack.Pop()
				res = MaxInt(res, i-start+1)
			} else {
				stack.Pop()
				t := stack.Pop()
				res = MaxInt(res, i-int(t))
				stack.Push(t)
			}
		}
	}
	return res
}
