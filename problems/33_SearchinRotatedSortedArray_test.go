package problems

import "testing"

func TestSearch(t *testing.T) {
	if search([]int{4, 5, 6, 7, 0, 1, 2}, 0) == 4 {
		t.Log("33. Search in Rotated Sorted Array case 1 succeeded.")
	} else {
		t.Error("33. Search in Rotated Sorted Array: case 1 failed.")
	}
	if search([]int{4, 5, 6, 7, 0, 1, 2}, 3) == -1 {
		t.Log("33. Search in Rotated Sorted Array case 2 succeeded.")
	} else {
		t.Error("33. Search in Rotated Sorted Array: case 2 failed.")
	}
	if search([]int{1}, 1) == 0 {
		t.Log("33. Search in Rotated Sorted Array case 2 succeeded.")
	} else {
		t.Error("33. Search in Rotated Sorted Array: case 2 failed.")
	}
}
