package problems

type Stack struct {
	dataSlice []uint8
}

func NewStack(size int) *Stack{
	return &Stack{make([]uint8, 0, size)}
}

func (stack *Stack) Push(val uint8) {
	stack.dataSlice = append(stack.dataSlice, val)
}

func (stack *Stack) Pop() uint8 {
	if len(stack.dataSlice) == 0 {
		return 0
	}
	var val = stack.dataSlice[len(stack.dataSlice) -1]
	stack.dataSlice = stack.dataSlice[:len(stack.dataSlice)-1]
	return val
}

func (stack *Stack) Size() int {
	return len(stack.dataSlice)
}

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.
//
//Example 1:
//
//	Input: "()"
//	Output: true
//
// Example 2:
//
//	Input: "()[]{}"
//	Output: true
//
// Example 3:
//
//	Input: "(]"
//	Output: false
//
// Example 4:
//
//	Input: "([)]"
//	Output: false
//
// Example 5:
//
//	Input: "{[]}"
//	Output: true
func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) % 2 == 1 {
		return false
	}
	var matcher = map[uint8]uint8{')': '(', ']': '[', '}': '{'}
	stack := NewStack(len(s))
	for _, char := range s {
		var val, ok = matcher[uint8(char)]
		if ok {
			pre := stack.Pop()
			if pre != val{
				return false
			}
		} else {
			stack.Push(uint8(char))
		}
	}
	return stack.Size() == 0
}
