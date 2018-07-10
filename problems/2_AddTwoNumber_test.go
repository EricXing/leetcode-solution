package problems

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	//
	var l1 = &ListNode{Val : 2}
	l1.Next = &ListNode{Val : 4}
	l1.Next.Next = &ListNode{Val:3}
	var l2 = &ListNode{Val : 5}
	l2.Next = &ListNode{Val : 6}
	l2.Next.Next = &ListNode{Val : 4}

	var result = AddTwoNumbers(l1, l2)
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Error("AddTwoNumbers with same length failed.")
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	var l1 = &ListNode{Val:1}
	l1.Next = &ListNode{Val:8}
	var l2 = &ListNode{Val:0}

	var result = AddTwoNumbers(l1, l2)
	if result.Val != 1 || result.Next.Val != 8 {
		t.Error("AddTwoNumbers with different lengths failed.")
	}
}
