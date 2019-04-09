package linkedlist

import (
	"sync"
)

func main() {

}

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (ll *LinkedList) Append(item interface{}) {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	node := Node{item, nil}
	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				last.next = &node
				break
			}
			last = last.next
		}
	}
	ll.size++

}

func (ll *LinkedList) Insert(i int, item interface{}) {
	ll.lock.RLock()
}

/*
Append(t) adds an Item t to the end of the linked list
Insert(i, t) adds an Item t at position i
RemoveAt(i) removes a node at position i
IndexOf() returns the position of the Item t
IsEmpty() returns true if the list is empty
Size() returns the linked list size
String() returns a string representation of the list
Head() returns the first node, so we can iterate on it
*/

func (ll *LinkedList) RemoveAt(i int) {

}

func (ll *LinkedList) Reverse() *Node {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	if ll.head == nil {
		return nil
	}
	cur := ll.head
	var prev, newHead *Node = nil, nil
	for cur != nil {
		next := cur.next
		cur.next = prev
		if next == nil {
			newHead = cur
		}
		prev = cur
		cur = next
	}
	return newHead
}
