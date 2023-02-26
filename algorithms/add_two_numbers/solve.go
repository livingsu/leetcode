package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	tail := prev
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		tail.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		tail = tail.Next
	}
	if carry == 1 {
		tail.Next = &ListNode{Val: 1}
	}
	return prev.Next
}
