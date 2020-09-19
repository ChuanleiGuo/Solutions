package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	cur := head
	size := 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	k %= size
	if size == k || k == 0 {
		return head
	}
	cur = head

	var prev *ListNode
	fast := head
	var endPrev *ListNode

	for i := 0; i < k; i++ {
		endPrev = fast
		fast = fast.Next
	}

	for fast != nil {
		prev = cur
		endPrev = fast
		cur = cur.Next
		fast = fast.Next
	}

	prev.Next = nil
	endPrev.Next = head

	return cur
}
