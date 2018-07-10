package problems

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var r1 = FindMedianSortedArrays([]int{1, 3}, []int{2})
	if r1 != 2.0 {
		t.Error("FindMedianSortedArrays: case 1 failed")
	}
	var r2 = FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	if r2 != 2.5 {
		t.Error("FindMedianSortedArrays: case 2 failed")
	}
	var r3 = FindMedianSortedArrays([]int{}, []int{2, 3})
	if r3 != 2.5 {
		t.Error("FindMedianSortedArrays: case 3 failed")
	}
	var r4 = FindMedianSortedArrays([]int{1}, []int{1})
	if r4 != 1.0 {
		t.Error("FindMedianSortedArrays: case 3 failed")
	}
}
