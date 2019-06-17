package main

import "fmt"

/*
反转链表
 */
func reverse(head *Node) {
	var prev, next *Node = nil, nil
	cur := head
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
}

/*
使用俩个指针反转链表
*/
func reverseWithTowPoints(head *Node) {
	var next *Node = nil
	cur := head

	for cur.next != nil {
		next = cur.next
		cur.next = next.next
		next.next = head
		head = next
	}
}

/*
反转链表的一部分
 */

/*
 每k个元素反转
Reverse a Linked List in groups of given size

 */

/*
以反转顺序打印链表，但不反转
 */
func printReverseList(head *Node) {
	if head == nil {
		return
	}
	printReverseList(head)
	fmt.Printf("%d  ", head.data)
}
