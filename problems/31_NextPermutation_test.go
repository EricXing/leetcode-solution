package problems

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	if reflect.DeepEqual(NextPermutation([]int{1, 2, 3}), []int{1, 3, 2}) {
		t.Log("31. Next Permutation: case 1 succeeded.")
	} else {
		t.Error("31. Next Permutation: case 1 failed.")
	}
	if reflect.DeepEqual(NextPermutation([]int{3, 2, 1}), []int{1, 2, 3}) {
		t.Log("31. Next Permutation: case 1 succeeded.")
	} else {
		t.Error("31. Next Permutation: case 1 failed.")
	}
	if reflect.DeepEqual(NextPermutation([]int{1, 1, 5}), []int{1, 5, 1}) {
		t.Log("31. Next Permutation: case 1 succeeded.")
	} else {
		t.Error("31. Next Permutation: case 1 failed.")
	}
}
