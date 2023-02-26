package _06__reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. iteration
func reverseList_iteration(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 2. recursion
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
