package _5__reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	prev := &ListNode{Next: head}
	prevTail, curHead, cur := prev, head, head
	i := 1
	for cur != nil {
		if i == k {
			nextHead := cur.Next
			cur.Next = nil
			reverseGroup(curHead)
			prevTail.Next = cur
			curHead.Next = nextHead

			prevTail = curHead
			curHead = nextHead
			cur = nextHead
			i = 1
		} else {
			cur = cur.Next
			i++
		}
	}
	return prev.Next
}

func reverseGroup(head *ListNode) {
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}
