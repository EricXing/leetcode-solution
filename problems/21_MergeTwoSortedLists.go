package problems

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
//	Input: 1->2->4, 1->3->4
//	Output: 1->1->2->3->4->4
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var p = &result
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if *p == nil {
			*p = &ListNode{}
		}
		if l1 == nil && l2 != nil {
			(*p).Val = l2.Val
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			(*p).Val = l1.Val
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				(*p).Val = l1.Val
				l1 = l1.Next
			} else {
				(*p).Val = l2.Val
				l2 = l2.Next
			}
		}
		p = &(*p).Next
	}
	return result
}
