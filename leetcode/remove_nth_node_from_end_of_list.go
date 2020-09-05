package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	h := dummy
	t := dummy
	for i := 0; i < n+1; i++ {
		t = t.Next
	}

	for t != nil {
		h = h.Next
		t = t.Next
	}
	h.Next = h.Next.Next
	return dummy.Next
}
