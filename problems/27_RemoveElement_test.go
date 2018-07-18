package problems

import "testing"

func TestRemoveElement(t *testing.T) {
	if RemoveElement([]int{3, 2, 2, 3}, 3) == 2 {
		t.Log("Remove Element: case 1 succeeded")
	} else {
		t.Error("Remove Element: case 1 failed")
	}
	if RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) == 5 {
		t.Log("Remove Element: case 2 succeeded")
	} else {
		t.Error("Remove Element: case 2 failed")
	}
}
