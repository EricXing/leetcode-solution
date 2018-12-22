package problems

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	if reflect.DeepEqual([]int{2, 2}, intersect([]int{1, 2, 2, 1}, []int{2, 2})) {
		t.Log("350. Intersection of Two Arrays II: case 1 succeeded.")
	} else {
		t.Error("350. Intersection of Two Arrays II: case 1 failed.")
	}
	if reflect.DeepEqual([]int{4, 9}, intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) {
		t.Log("350. Intersection of Two Arrays II: case 2 succeeded.")
	} else {
		t.Error("350. Intersection of Two Arrays II: case 2 failed.")
	}
}
