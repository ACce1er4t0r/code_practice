package swapPairs

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    cur, pre := head, head.Next
    cur.Next = swapPairs(pre.Next)
    pre.Next = cur
    return pre
}