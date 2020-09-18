package leetcode

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root

	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
		} else {
			next := cur.Next
			firstChild := flatten(cur.Child)
			lastChild := firstChild
			for lastChild.Next != nil {
				lastChild = lastChild.Next
			}
			lastChild.Next = next
			if next != nil {
				next.Prev = lastChild
			}
			cur.Next = firstChild
			firstChild.Prev = cur
			cur.Child = nil
			cur = next
		}
	}
	return root
}
