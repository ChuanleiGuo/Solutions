package leetcode

type node struct {
	val  int
	next *node
}

type MyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func Constructor() MyLinkedList {
	n := &node{}
	head := n
	tail := head
	length := 0
	l := MyLinkedList{
		head:   head,
		tail:   tail,
		length: length,
	}
	return l
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.length {
		return -1
	}
	ptr := l.head.next
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	return ptr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	n := &node{
		val: val,
	}
	n.next = l.head.next
	l.head.next = n

	if n.next == nil {
		l.tail = n
	}
	l.length++
}

func (l *MyLinkedList) AddAtTail(val int) {
	n := &node{
		val: val,
	}
	n.next = l.tail.next
	l.tail.next = n

	l.tail = n
	l.length++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.length {
		return
	}
	if index == 0 {
		l.AddAtHead(val)
	} else if index == l.length {
		l.AddAtTail(val)
	} else {
		pre := l.head
		cur := l.head.next

		for i := 0; i < index; i++ {
			pre = pre.next
			cur = cur.next
		}

		n := &node{val: val}
		n.next = cur
		pre.next = n
		l.length++
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.length {
		return
	}
	pre := l.head
	cur := l.head.next
	for i := 0; i < index; i++ {
		pre = pre.next
		cur = cur.next
	}

	pre.next = cur.next
	cur.next = nil
	if index == l.length-1 {
		l.tail = pre
	}
	l.length--
}
