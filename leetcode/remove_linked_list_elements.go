package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dummy
	curr := head
	for curr != nil {
		if curr.Val != val {
			prev = prev.Next
			curr = curr.Next
			continue
		}
		prev.Next = curr.Next
		curr = prev.Next
	}
	return dummy.Next
}
