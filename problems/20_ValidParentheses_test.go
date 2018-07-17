package problems

import "testing"

func TestIsValid(t *testing.T) {
	if !IsValid("()"){
		t.Error("Valid Parentheses: case 1 faled.")
	} else {
		t.Log("Valid Parentheses: case 1 succeeded.")
	}
	if !IsValid("()[]{}"){
		t.Error("Valid Parentheses: case 2 faled.")
	} else {
		t.Log("Valid Parentheses: case 2 succeeded.")
	}
	if IsValid("(]"){
		t.Error("Valid Parentheses: case 3 faled.")
	} else {
		t.Log("Valid Parentheses: case 3 succeeded.")
	}
	if IsValid("([)]"){
		t.Error("Valid Parentheses: case 4 faled.")
	} else {
		t.Log("Valid Parentheses: case 4 succeeded.")
	}
	if !IsValid("{[]}"){
		t.Error("Valid Parentheses: case 5 faled.")
	} else {
		t.Log("Valid Parentheses: case 5 succeeded.")
	}
	if IsValid("["){
		t.Error("Valid Parentheses: case 6 faled.")
	} else {
		t.Log("Valid Parentheses: case 6 succeeded.")
	}
	if IsValid("(("){
		t.Error("Valid Parentheses: case 7 faled.")
	} else {
		t.Log("Valid Parentheses: case 7 succeeded.")
	}
}
