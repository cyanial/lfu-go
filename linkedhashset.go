package lfu

type LinkedHashSet struct {
	key2node map[int]*DNode
	cache    DoubleLinkedList
}

func (l *LinkedHashSet) addRecently(key int) {
	x := &DNode{key, nil, nil}
	l.cache.addLast(x)
	l.key2node[key] = x
}

func (l *LinkedHashSet) deleteKey(key int) {
	x := l.key2node[key]
	remove(x)
	delete(l.key2node, key)
}

func (l *LinkedHashSet) add(key int) {
	if _, ok := l.key2node[key]; ok {
		l.deleteKey(key)
		l.addRecently(key)
		return
	}
	l.addRecently(key)
}

func createdLinkedHashSet() *LinkedHashSet {
	head := &DNode{0, nil, nil}
	tail := &DNode{0, nil, nil}
	head.next = tail
	tail.prev = head
	doubledLinkedList := DoubleLinkedList{head, tail}
	key2node := map[int]*DNode{}
	return &LinkedHashSet{
		key2node: key2node,
		cache:    doubledLinkedList,
	}
}

func (l *LinkedHashSet) isEmpty() bool {
	return len(l.key2node) == 0
}
