package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = head
	for slow != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
