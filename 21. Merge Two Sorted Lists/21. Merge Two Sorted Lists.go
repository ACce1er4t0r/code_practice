package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	res := head
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 == nil {
		head.Next = l2
	} else if l2 == nil {
		head.Next = l1
	}
	return res.Next
}
