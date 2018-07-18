package problems

import "testing"

func TestSearchInsert(t *testing.T) {
	if SearchInsert([]int{1, 3, 5, 6}, 5) == 2 {
		t.Log("Search Insert Position: case 1 succeeded.")
	} else {
		t.Error("Search Insert Position: case 1 failed")
	}
	if SearchInsert([]int{1, 3, 5, 6}, 2) == 1 {
		t.Log("Search Insert Position: case 2 succeeded.")
	} else {
		t.Error("Search Insert Position: case 2 failed")
	}
	if SearchInsert([]int{1, 3, 5, 6}, 7) == 4 {
		t.Log("Search Insert Position: case 3 succeeded.")
	} else {
		t.Error("Search Insert Position: case 3 failed")
	}
	if SearchInsert([]int{1, 3, 5, 6}, 0) == 0 {
		t.Log("Search Insert Position: case 4 succeeded.")
	} else {
		t.Error("Search Insert Position: case 4 failed")
	}
}
