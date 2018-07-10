package problems

import (
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (listNode *ListNode) String() string {
	var s = "("
	for i := listNode; i!=nil; i = i.Next{
		s = s + strconv.Itoa(i.Val) + ", "
	}
	s += ")"
	return s
}

//You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single
// digit. Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example
//	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//	Output: 7 -> 0 -> 8
//	Explanation: 342 + 465 = 807.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{Val:0}
	var count = 0
	for i, j := l1, l2 ; i!=nil || j !=nil; {
		var index = &result
		for z:=0; z<count ; z++ {
			if (*index).Next == nil {
				(*index).Next = &ListNode{Val:0}
			}
			index = &(*index).Next
		}
		var x int
		if i != nil {
			x = i.Val
		}
		var y int
		if j != nil {
			y = j.Val
		}
		(*index).Val = (*index).Val + x + y
		if (*index).Val >= 10 {
			(*index).Val -= 10
			(*index).Next = &ListNode{Val:1}
		}
		count += 1
		if i != nil {
			i = i.Next
		}
		if j != nil {
			j = j.Next
		}
	}
	return result
}


