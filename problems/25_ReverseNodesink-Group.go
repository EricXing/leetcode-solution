package problems

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list.
// If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
// Example:
//
//	Given this linked list: 1->2->3->4->5
//
//	For k = 2, you should return: 2->1->4->3->5
//
//	For k = 3, you should return: 3->2->1->4->5
//
// Note:
//
//	Only constant extra memory is allowed.
//	You may not alter the values in the list's nodes, only nodes itself may be changed.
func ReverseKGroup(head *ListNode, k int) *ListNode {
	size := 0
	temp := &head
	for *temp != nil {
		size += 1
		temp = &(*temp).Next
	}

	root := &ListNode{Next: head}
	res := root
	for size >= k {
		for i := 0; i < k-1; i += 1 {
			node := root.Next
			root.Next = head.Next
			head.Next = root.Next.Next
			root.Next.Next = node
		}
		root = head
		head = head.Next
		size -= k
	}

	return res.Next
}
