package problems

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	re := MergeTwoLists(l1, l2)

	if re.Val == 1 && re.Next.Val == 1 && re.Next.Next.Val == 2 && re.Next.Next.Next.Val == 3 && re.Next.Next.Next.Next.Val == 4 &&
		re.Next.Next.Next.Next.Next.Val == 4 && re.Next.Next.Next.Next.Next.Next == nil {
		t.Log("21. Merge Two Sorted Lists: case 1 succeeded.")
	} else {
		t.Error("21. Merge Two Sorted Lists: case 1 failed.")
	}
}
