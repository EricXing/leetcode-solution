package problems

import "testing"

func TestMaxArea(t *testing.T) {
	if MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49 {
		t.Log("Container With Most Water: case 1 succeeded.")
	} else {
		t.Error("Container With Most Water: case 1 failed.")
	}
	if MaxArea([]int{2, 3, 4, 5, 18, 17, 6}) == 17 {
		t.Log("Container With Most Water: case 1 succeeded.")
	} else {
		t.Error("Container With Most Water: case 1 failed.")
	}
}
