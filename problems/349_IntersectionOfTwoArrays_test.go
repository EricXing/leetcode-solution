package problems

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	if reflect.DeepEqual([]int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2})) {
		t.Log("349. Intersection of Two Arrays: case 1 succeeded.")
	} else {
		t.Error("349. Intersection of Two Arrays: case 1 failed.")
	}
	if reflect.DeepEqual([]int{4, 9}, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) {
		t.Log("349. Intersection of Two Arrays: case 2 succeeded.")
	} else {
		t.Error("349. Intersection of Two Arrays: case 2 failed.")
	}
}
