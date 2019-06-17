package main

/*
交换linkedlist中的俩个节点
1. 可以通过prev和cur来做，分为俩次查询： 报错prev
prevX.next=curY
prevY.next =curX
temp = curY.next
curY.next=curX.next
curX.next=temp

2. 可以通过一次查询，交换节点，然后交换节点的next节点，这种方法只需要保存curx，和cury即可
swap(curx,cury)
swap(curx.next,cury.next)

swap(a,b node){
	tmp = a
	a=b
	b=tmp
}
 */


/*
反转linkedlist
1. 反转linkedlist： 需要保存三个节点： prev，cur next三个节点，cur=head,prev,next=nil
for cur != nil {
	next = cur.next
	cur.next=prev
	prev=cur
	cur=next
}
 */


/**
 * 160. Intersection of Two Linked Lists 相交链表
 *
 * 编写一个程序，找到两个单链表相交的起始节点。
 *
 * 如下面的两个链表,在节点 c1 开始相交:
 *
 * 示例 1：
 *       4 -> 1
 *              -> 8 -> 4 -> 5
 *  5 -> 0 -> 1
 *
 *      输入：
 *          intersectVal = 8,
 *          listA = [4,1,8,4,5],
 *          listB = [5,0,1,8,4,5],
 *          skipA = 2,
 *          skipB = 3
 *      输出：Reference of the node with value = 8
 *      解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。
 *           从各自的表头开始算起，链表 A 为 [4, 1, 8, 4, 5]，链表 B
 *           为 [5, 0, 1, 8, 4, 5]。在 A 中，相交节点前有 2 个节点；
 *           在 B 中，相交节点前有 3 个节点。
 *
 * 2019-04-11 11:26 PM
利用求回环的来解： 如果有相同的节点，只需要x+y =y+x个length即可找到

for curX !=nil  && curY !=nil {
	if curX == nil {
		curX=headY
	}else {
		curX=curX.next
	}
	if curY ==nil {
		curY=headX
	}else {
		curY=curY.next
	}
}
 **/

/*
单链表的quickSort

 */


 /*
 单链表的mergesort
  */