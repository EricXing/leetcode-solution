package problems

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	if reflect.DeepEqual(ThreeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, 0, 1}, {-1, -1, 2}}) {
		t.Log("15. 3Sum: case 1 succeeded.")
	} else {
		t.Error("15. 3Sum case 1 failed.")
	}
	if reflect.DeepEqual(ThreeSum([]int{0, 0}), [][]int{}) {
		t.Log("15. 3Sum: case 2 succeeded.")
	} else {
		t.Error("15. 3Sum case 2 failed.")
	}
	if reflect.DeepEqual(ThreeSum([]int{-1, 0, 1, 0}), [][]int{{-1, 0, 1}}) {
		t.Log("15. 3Sum: case 2 succeeded.")
	} else {
		t.Error("15. 3Sum case 2 failed.")
	}
	if reflect.DeepEqual(ThreeSum([]int{1, 2, -2, -1}), [][]int{}) {
		t.Log("15. 3Sum: case 3 succeeded.")
	} else {
		t.Error("15. 3Sum case 3 failed.")
	}
}
