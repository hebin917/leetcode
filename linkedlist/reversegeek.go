package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

/*
使用俩个指针来反转linkedlist
 Function to reverse the linked list using 2 pointers
//void reverse(struct Node** head_ref)
//{
//struct Node* current = *head_ref;
//struct Node* next;
//while (current->next != NULL) {
//next = current->next;
//current->next = next->next;
//next->next = (*head_ref);
//*head_ref = next;
//}
//}
*/
func x() {

}

/*
Reverse a Linked List in groups of given size | Set 1   反转一组given size
 */
func reverse(head *Node, k int) *Node {

	cur := head
	var next, prev *Node
	i := 0
	for cur != nil && i < k {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
		i++
	}
	if next != nil {
		head.next = reverse(next, k)
	}
	return prev
}

//方法二： 使用stack来完成
/*
1. 先pushk个元素，并且保持prev
2.从prev开始重复更新反转下一个k
 */
func reverseWithStack(head *Node, k int) *Node {
	cur := head
	var prev *Node =nil
	s := stack.New()
	for cur != nil {
		i := 0
		for cur != nil && i < k {
			s.Push(cur)
			cur = cur.next
			i++
		}
		for s.Len() > 0 {
			if prev == nil {
				prev = s.Pop().(*Node)
				head = prev
			} else {
				prev.next = s.Pop().(*Node)
				prev = prev.next
			}
		}

	}
	prev.next = nil
	return head
}

func main() {
	head := NewNode()
	//head = reverse(head, 5)
	head=reverseWithStack(head, 5)
	for head != nil {
		fmt.Printf("%d  ", head.data)
		head = head.next
	}

}

func NewNode() *Node {
	node := Node{1, &Node{2, &Node{
		3, &Node{
			4, &Node{5,
				&Node{6, &Node{7, nil}},
			},
		},
	}}}
	return &node
}

type Node struct {
	data interface{}
	next *Node
}
