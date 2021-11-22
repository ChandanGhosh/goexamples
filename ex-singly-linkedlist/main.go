package main

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head  *Node
	count int
}

func (l *LinkedList) InsertFirst(data int) {
	node := &Node{val: data}
	node.next = l.head
	l.head = node
	l.count++
}

func (l *LinkedList) Find(val int) bool {
	item := l.head
	for item != nil {
		if item.val == val {
			return true
		} else {
			item = item.next
		}
	}
	return false
}

func (l *LinkedList) DeleteAt(idx int) {
	if idx > l.count-1 {
		return
	}
	if idx == 0 {
		l.head = l.head.next
		l.count--
	} else {
		var tempIdx = 0
		var node = l.head
		for tempIdx < idx-1 {
			node = node.next
			tempIdx++
		}
		nodeToDelete := node.next
		node.next = nodeToDelete.next
		nodeToDelete = nil
	}
}

func (l *LinkedList) DeleteFirst(val int) bool {
	item := l.head
	for item != nil {
		if item.val == val {
			l.head = item.next
			item = nil
			l.count--
			return true
		} else {
			return false
		}
	}
	return false
}

func (l *LinkedList) print() {
	item := l.head
	for item != nil {
		println(item.val)
		item = item.next
	}
}

func main() {
	var linkedList LinkedList
	linkedList.InsertFirst(10)
	linkedList.InsertFirst(20)
	linkedList.InsertFirst(30)
	linkedList.InsertFirst(40)

	linkedList.print()

	println(linkedList.Find(40))
	linkedList.DeleteAt(5)
	linkedList.print()
}
