package problems

import "fmt"

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// 
// For example, given n = 3, a solution set is:
// 
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]
func GenerateParenthesis(n int) []string {
	result := make([]string, 0)
	if n == 0{
		result = append(result, "")
		return result
	}
	for i := 0; i< n;i++ {
		for _, left:= range GenerateParenthesis(i) {
			for _, right := range GenerateParenthesis(n-1-i){
				result = append(result, fmt.Sprintf("(%s)%s", left, right))
			}
		}
	}
	return result
}
