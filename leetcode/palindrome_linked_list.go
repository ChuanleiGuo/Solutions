package leetcode

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverse := func(h *ListNode) *ListNode {
		var prev *ListNode
		curr := h

		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		return prev
	}

	fast = head
	slow = reverse(slow)

	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
