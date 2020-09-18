package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead
	cur1 := l1
	cur2 := l2

	carry := 0
	for cur1 != nil || cur2 != nil {

		x := 0
		if cur1 != nil {
			x = cur1.Val
		}

		y := 0
		if cur2 != nil {
			y = cur2.Val
		}

		sum := x + y + carry
		val := sum % 10
		carry = sum / 10

		tail.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		tail = tail.Next

		if cur1 != nil {
			cur1 = cur1.Next
		}

		if cur2 != nil {
			cur2 = cur2.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry, Next: nil}
	}

	return dummyHead.Next
}
