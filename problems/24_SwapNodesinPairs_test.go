package problems

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	input1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	res1 := SwapPairs(input1)
	if reflect.DeepEqual(res1, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}}) {
		t.Log("24. Swap Nodes in Pairs: case 1 succeeded.")
	} else {
		t.Error("24. Swap Nodes in Pairs: case 1 failed.")
	}
}
