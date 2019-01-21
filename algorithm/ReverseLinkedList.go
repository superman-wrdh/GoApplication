package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Reverse Linked List
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next

	head.Next = nil

	rest := reverseList(second)

	second.Next = head

	return rest
}
