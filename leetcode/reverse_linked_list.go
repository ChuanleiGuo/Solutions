package leetcode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverserListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverserListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
