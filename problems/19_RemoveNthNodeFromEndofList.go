package problems

// Given a linked list, remove the n-th node from the end of list and return its head.
//
// Example:
//
// Given linked list: 1->2->3->4->5, and n = 2.
//
//	After removing the second node from the end, the linked list becomes 1->2->3->5.
//
// Note:
//
// Given n will always be valid.
//
// Follow up:
//
// Could you do this in one pass?
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	nodeMap := make(map[int]**ListNode)
	count := 0
	curr := &head
	for ; (*curr) != nil; curr, count = &((*curr).Next), count+1 {
		nodeMap[count] = curr
	}
	length := len(nodeMap)
	firstIdx := length - n - 1
	lastIdx := length - n + 1
	if firstIdx < 0{
		if length == 1 {
			return nil
		} else {
			return *(nodeMap[1])
		}
	}
	if lastIdx == length {
		(*nodeMap[length-2]).Next = nil
	} else {
		(*nodeMap[firstIdx]).Next = *(nodeMap[lastIdx])
	}
	return head
}
