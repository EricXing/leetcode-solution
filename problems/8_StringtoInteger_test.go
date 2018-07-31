package problems

import "testing"

func TestMyAtoi(t *testing.T) {
	if MyAtoi("42") == 42 {
		t.Log("String to Integer: case 1 succeeded.")
	} else {
		t.Error("String to Integer: case 1 failed.")
	}
	if MyAtoi("  -42") == -42 {
		t.Log("String to Integer: case 2 succeeded.")
	} else {
		t.Error("String to Integer: case 2 failed.")
	}
	if MyAtoi("4193 with words") == 4193 {
		t.Log("String to Integer: case 3 succeeded.")
	} else {
		t.Error("String to Integer: case 3 failed.")
	}
	if MyAtoi("words and 987") == 0 {
		t.Log("String to Integer: case 4 succeeded.")
	} else {
		t.Error("String to Integer: case 4 failed.")
	}
	if MyAtoi("-91283472332") == -2147483648 {
		t.Log("String to Integer: case 5 succeeded.")
	} else {
		t.Error("String to Integer: case 5 failed.")
	}
	if MyAtoi("+1") == 1 {
		t.Log("String to Integer: case 6 succeeded.")
	} else {
		t.Error("String to Integer: case 6 failed.")
	}
	if MyAtoi("   +0 123") == 0 {
		t.Log("String to Integer: case 7 succeeded.")
	} else {
		t.Error("String to Integer: case 7 failed.")
	}
	if MyAtoi("-5-") == -5 {
		t.Log("String to Integer: case 8 succeeded.")
	} else {
		t.Error("String to Integer: case 8 failed.")
	}
}
