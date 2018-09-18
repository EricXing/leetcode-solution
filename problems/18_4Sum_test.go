package problems

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	if reflect.DeepEqual(FourSum([]int{1, 0, -1, 0, -2, 2}, 0), [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}) {
		t.Log("18. 4Sum: case 1 succeeded.")
	} else {
		t.Error("18. 4Sum: case 1 failed.")
	}
}
