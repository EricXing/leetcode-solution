package problems

// Given a linked list, swap every two adjacent nodes and return its head.
//
// Example:
//
// 	Given 1->2->3->4, you should return the list as 2->1->4->3.
//
// Note:
//
// Your algorithm should use only constant extra space.
// You may not modify the values in the list's nodes, only nodes itself may be changed.
func SwapPairs(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	root := res
	curr := head
	for curr != nil && curr.Next != nil {
		node := root.Next
		root.Next = curr.Next
		curr.Next = root.Next.Next
		root.Next.Next = node
		root = curr
		curr = curr.Next
	}
	return res.Next
}
