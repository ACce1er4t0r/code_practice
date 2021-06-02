package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{Val: 0, Next: nil}
	head := res
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum < 10 {
			res.Val = sum
			carry = 0
		} else {
			res.Val = sum % 10
			carry = sum / 10
		}
		if l1 != nil || l2 != nil || carry != 0 {
			res.Next = &ListNode{Val: 0, Next: nil}
			res = res.Next
		}
	}
	return head
}
