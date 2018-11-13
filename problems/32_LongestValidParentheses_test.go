package problems

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	if LongestValidParentheses("(()") == 2 {
		t.Log("32. Longest Valid Parentheses: case 1 succeeded.")
	} else {
		t.Error("32. Longest Valid Parentheses: case 1 failed.")
	}
	if LongestValidParentheses(")()())") == 4 {
		t.Log("32. Longest Valid Parentheses: case 2 succeeded.")
	} else {
		t.Error("32. Longest Valid Parentheses: case 2 failed.")
	}
	if LongestValidParentheses("()(())") == 6 {
		t.Log("32. Longest Valid Parentheses: case 3 succeeded.")
	} else {
		t.Error("32. Longest Valid Parentheses: case 3 failed.")
	}
	if LongestValidParentheses("))))((()((") == 2 {
		t.Log("32. Longest Valid Parentheses: case 4 succeeded.")
	} else {
		t.Error("32. Longest Valid Parentheses: case 4 failed.")
	}
	if LongestValidParentheses("(()()") == 4 {
		t.Log("32. Longest Valid Parentheses: case 5 succeeded.")
	} else {
		t.Error("32. Longest Valid Parentheses: case 5 failed.")
	}
}
