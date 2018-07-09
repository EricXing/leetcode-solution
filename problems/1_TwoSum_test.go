package problems

import "testing"

func TestTwoSum(t *testing.T){
	var param = []int{2, 7, 11, 5}
	var result = TwoSum(param, 9)
	if result[0] != 0 || result[1] != 1 {
		t.Error("TwoSum failed")
	}
}
