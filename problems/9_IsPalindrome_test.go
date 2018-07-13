package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	var num1 = 121
	if !IsPalindrome(num1) {
		t.Error("case 1 failed")
	}
	var num2 = -121
	if IsPalindrome(num2) {
		t.Error("case 2 failed")
	}
	var num3 = 10
	if IsPalindrome(num3) {
		t.Error("case 3 failed")
	}
	var num4 = 123
	if IsPalindrome(num4) {
		t.Error("case 4 failed")
	}
}
