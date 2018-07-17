package problems

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	if RemoveDuplicates([]int{1, 1, 2}) == 2{
		t.Log("Remove Duplicates from Sorted Array: case 1 succeeded.")
	} else {
		t.Error("Remove Duplicates from Sorted Array case 1 failed.")
	}
	if RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) == 5{
		t.Log("Remove Duplicates from Sorted Array: case 2 succeeded.")
	} else {
		t.Error("Remove Duplicates from Sorted Array case 2 failed.")
	}
}
