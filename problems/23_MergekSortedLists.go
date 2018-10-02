package problems

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
// Example:
//
//	Input:
//	[
//	  1->4->5,
//	  1->3->4,
//	  2->6
//	]
//	Output: 1->1->2->3->4->4->5->6
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var result = MergeTwoLists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		result = MergeTwoLists(result, lists[i])
	}
	return result
}
