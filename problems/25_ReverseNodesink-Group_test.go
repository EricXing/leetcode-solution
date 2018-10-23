package problems

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	input1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	input2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	res1 := ReverseKGroup(input1, 2)
	if reflect.DeepEqual(res1, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}){
		t.Log("25. Reverse Nodes in k-Group: case 1 succeeded.")
	} else {
		t.Error("25. Reverse Nodes in k-Group: case 1 failed.")
	}
	res2 := ReverseKGroup(input2, 3)
	if reflect.DeepEqual(res2, &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}){
		t.Log("25. Reverse Nodes in k-Group: case 1 succeeded.")
	} else {
		t.Error("25. Reverse Nodes in k-Group: case 1 failed.")
	}
}
