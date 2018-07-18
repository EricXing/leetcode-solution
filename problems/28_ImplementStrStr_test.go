package problems

import "testing"

func TestStrStr(t *testing.T) {
	if StrStr("hello", "ll") == 2 {
		t.Log("Implement strStr: case 1 succeeded.")
	} else {
		t.Error("Implement strStr: case 1 failed.")
	}
	if StrStr("aaaaa", "bba") == -1 {
		t.Log("Implement strStr: case 2 succeeded.")
	} else {
		t.Error("Implement strStr: case 2 failed.")
	}
	if StrStr("a", "a") == 0 {
		t.Log("Implement strStr: case 3 succeeded.")
	} else {
		t.Error("Implement strStr: case 3 failed.")
	}
	if StrStr("mississippi", "pi") == 9 {
		t.Log("Implement strStr: case 4 succeeded.")
	} else {
		t.Error("Implement strStr: case 4 failed.")
	}
}