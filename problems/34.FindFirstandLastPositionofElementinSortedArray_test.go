package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	if reflect.DeepEqual(searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4}) {
		t.Log("34. Find First and Last Position of Element in Sorted Array case 1 succeeded.")
	} else {
		t.Error("34. Find First and Last Position of Element in Sorted Array case 1 failed.")
	}
	if reflect.DeepEqual(searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1}) {
		t.Log("34. Find First and Last Position of Element in Sorted Array case 2 succeeded.")
	} else {
		t.Error("34. Find First and Last Position of Element in Sorted Array case 2 failed.")
	}
	if reflect.DeepEqual(searchRange([]int{1, 1, 2, 3, 3}, 2), []int{2, 2}) {
		t.Log("34. Find First and Last Position of Element in Sorted Array case 3 succeeded.")
	} else {
		t.Error("34. Find First and Last Position of Element in Sorted Array case 3 failed.")
	}
	if reflect.DeepEqual(searchRange([]int{1, 1, 2, 3, 3}, 2), []int{2, 2}) {
		t.Log("34. Find First and Last Position of Element in Sorted Array case 4 succeeded.")
	} else {
		t.Error("34. Find First and Last Position of Element in Sorted Array case 4 failed.")
	}
	if reflect.DeepEqual(searchRange([]int{1, 3}, 1), []int{0, 0}) {
		t.Log("34. Find First and Last Position of Element in Sorted Array case 5 succeeded.")
	} else {
		t.Error("34. Find First and Last Position of Element in Sorted Array case 5 failed.")
	}
	fmt.Println(searchRange([]int{1, 4}, 4))
	if reflect.DeepEqual(searchRange([]int{1, 4}, 4), []int{1, 1}) {
		t.Log("34. Find First and Last Position of Element in Sorted Array case 6 succeeded.")
	} else {
		t.Error("34. Find First and Last Position of Element in Sorted Array case 6 failed.")
	}
}
