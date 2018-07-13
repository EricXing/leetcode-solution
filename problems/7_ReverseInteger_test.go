package problems

import "testing"

func TestReverse(t *testing.T) {
	if Reverse(123) != 321 {
		t.Error("Reverse Integer: case 1 failed")
	}
	if Reverse(-123) != -321 {
		t.Error("Reverse Integer: case 2 failed")
	}
	if Reverse(120) != 21 {
		t.Error("Reverse Integer: case 3 failed")
	}
	if Reverse(1534236469) != 0 {
		t.Error("Reverse Integer: case 3 failed")
	}
}
