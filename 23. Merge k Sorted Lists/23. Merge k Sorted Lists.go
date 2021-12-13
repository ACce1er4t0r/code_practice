package mergeKLists

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists1(lists []*ListNode) *ListNode {
    if len(lists) == 0{
		return nil
	}else if len(lists) == 1{
		return lists[0]
	}
    begin, end := 0, len(lists)-1

	for begin < end {
		mid := (begin + end - 1) / 2
		for i := 0; i <= mid; i++ {
			lists[i] = mergeTwoLists1(lists[i], lists[end-i])
		}
		end = (begin + end) / 2
	}
	return lists[0]
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil{
		return l2
	}else if l2 == nil{
		return l1
	}
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

func mergeKLists2(lists []*ListNode) *ListNode {
    if len(lists) == 0{
		return nil
	}else if len(lists) == 1{
		return lists[0]
	}
    var ret = lists[0]

	for i := 1; i < len(lists); i++{
		ret = mergeTwoLists2(ret, lists[i])
	}
    return ret
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil{
		return l2
	}else if l2 == nil{
		return l1
	}
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