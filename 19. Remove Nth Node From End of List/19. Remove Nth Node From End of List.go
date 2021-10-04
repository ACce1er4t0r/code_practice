package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
		if p1 == nil {
			return head.Next
		}
	}
	p2 := head
	for {
		p1 = p1.Next
		if p1 == nil {
			break
		}
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return head
}
