package problems

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	if reflect.DeepEqual([]int{1, 2}, twoSumII([]int{2, 7, 11, 15}, 9)) {
		t.Log("167. Two Sum II - Input array is sorted: case 1 succeeded.")
	} else {
		t.Error("167. Two Sum II - Input array is sorted: case 1 failed.")
	}
	if reflect.DeepEqual([]int{1, 3}, twoSumII([]int{2, 3, 4}, 6)) {
		t.Log("167. Two Sum II - Input array is sorted: case 2 succeeded.")
	} else {
		t.Error("167. Two Sum II - Input array is sorted: case 2 failed.")
	}
	if reflect.DeepEqual([]int{3, 6}, twoSumII([]int{3, 24, 50, 79, 88, 150, 345}, 200)) {
		t.Log("167. Two Sum II - Input array is sorted: case 3 succeeded.")
	} else {
		t.Error("167. Two Sum II - Input array is sorted: case 3 failed.")
	}

}
