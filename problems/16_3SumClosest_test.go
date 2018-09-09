package problems

import "testing"

func TestThreeSumClosest(t *testing.T) {
	if ThreeSumClosest([]int{-1, 2, 1, -4}, 1) == 2 {
		t.Log("16. 3Sum Closest: case 1 succeeded.")
	} else {
		t.Error("16. 3Sum Closest: case 1 failed.")
	}
}
