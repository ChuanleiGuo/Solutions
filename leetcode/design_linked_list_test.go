package leetcode

import "testing"

func TestMyLinkedList(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	v1 := l.Get(1)
	l.DeleteAtIndex(1)
	v2 := l.Get(1)

	if v1 != 2 {
		t.Errorf("error")
	}
	if v2 != 3 {
		t.Errorf("error")
	}
}
