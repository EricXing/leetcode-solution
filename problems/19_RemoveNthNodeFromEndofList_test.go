package problems

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := RemoveNthFromEnd(list, 2)
	if result.Val != 1 || result.Next.Val != 2 || result.Next.Next.Val != 3 || result.Next.Next.Next.Val != 5 {
		t.Error("19. Remove Nth Node From End of List: case 1 failed.")
	} else {
		t.Log("19. Remove Nth Node From End of List: case 1 succeeded.")
	}
	result2 :=RemoveNthFromEnd(&ListNode{Val: 1}, 1)
	if result2 == nil {
		t.Log("19. Remove Nth Node From End of List: case 2 succeeded.")
	} else {
		t.Error("19. Remove Nth Node From End of List: case 2 failed.")
	}
}
