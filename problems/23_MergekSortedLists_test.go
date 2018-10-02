package problems

import "testing"

func TestMergeKLists(t *testing.T) {
	var lists = []*ListNode{{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}
	result := MergeKLists(lists)
	if result.Val == 1 && result.Next.Val == 1 && result.Next.Next.Val == 2 &&
		result.Next.Next.Next.Val == 3 && result.Next.Next.Next.Next.Val == 4 &&
		result.Next.Next.Next.Next.Next.Val == 4 && result.Next.Next.Next.Next.Next.Next.Val == 5 &&
		result.Next.Next.Next.Next.Next.Next.Next.Val == 6 && result.Next.Next.Next.Next.Next.Next.Next.Next == nil {
			t.Log("23. Merge k Sorted Lists: case 1 succeeded.")
	} else {
		t.Log("23. Merge k Sorted Lists: case 1 failed.")
	}
}
