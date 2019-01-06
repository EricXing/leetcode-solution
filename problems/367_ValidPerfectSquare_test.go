package problems

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	if isPerfectSquare(16) {
		t.Log("367. Valid Perfect Square: case 1 succeeded.")
	} else {
		t.Log("367. Valid Perfect Square: case 1 failed.")
	}
	if !isPerfectSquare(14) {
		t.Log("367. Valid Perfect Square: case 2 succeeded.")
	} else {
		t.Log("367. Valid Perfect Square: case 2 failed.")
	}
}
