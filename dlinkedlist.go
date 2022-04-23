package lfu

type DoubleLinkedList struct {
	head *DNode
	tail *DNode
}

type DNode struct {
	key  int
	next *DNode
	prev *DNode
}

func remove(x *DNode) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

func (d *DoubleLinkedList) addLast(x *DNode) {
	x.prev = d.tail.prev
	x.next = d.tail
	d.tail.prev.next = x
	d.tail.prev = x
}

func (d *DoubleLinkedList) removeFirst() *DNode {
	if d.head.next == d.tail {
		return nil
	}
	first := d.head.next
	remove(first)
	return first
}
