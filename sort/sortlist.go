package main

import (
	"fmt"
	)


func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		preindex := i - 1
		current := nums[i]
		for preindex >= 0 && nums[preindex] > current {
			nums[preindex+1] = nums[preindex]
			preindex--
		}
		nums[preindex+1] = current
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
每次找到max，然后替换最大pre
 */
//func insertionSortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	fast := head
//	slow :=head
//	var pre *ListNode = nil
//	for fast !=nil && fast.Next!=nil {
//		pre=slow
//		pre=slow.Next
//		fast=fast.Next.Next
//	}
//
//	l := insertionSortList(head)
//	r := insertionSortList(slow)
//	return merge(l,r)
//}


/*
对链表进行排序，是对链表中的数据进行排序
先将链表中的数据取出，然后排序
然后覆盖链表中的值即可。
 */
func sortList(head *ListNode) *ListNode {
	vals := []int{}
	old := head
	maxlen := 0
	for old != nil {
		vals=append(vals,old.Val)
		old=old.Next
		maxlen++
	}

	fmt.Printf("maxlen.......%d \n",maxlen)
	// 对vals排序
	sort(vals)
	for _,v:= range vals {
		fmt.Println(v)
	}
	fmt.Println("--------------")
	//对排序后的数据，然后拼接成ListNode的链表
	//head.Val=vals[0]
	old=head
	for _,v :=range(vals) {
		fmt.Println(v)
		head.Val=v
		head=head.Next
	}
	fmt.Println("================")
	return old
}

/*
使用插入排序对ListNode进行排序
 */
func insertionSortList(head *ListNode) *ListNode {

	return nil
}

func sort(nums []int)  {
	for i:=1;i<len(nums);i++ {
		preindex := i - 1
		cur := nums[i]
		for preindex >= 0 && cur < nums[preindex] {
			nums[preindex+1] = nums[preindex]
			preindex--
		}
		nums[preindex+1] = cur
	}
}

func main() {
	//nums := []int{-1,5,3,4,0}
	//insertSort(nums)
	//for i :=0;i<len(nums);i++ {
	//	fmt.Println(nums[i])
	//}
	//fmt.Println("=====")
	//nums=[]int{100, 10, 40, 9, 20, 50, 21, 33}
	//insertSort(nums)
	//for i :=0;i<len(nums);i++ {
	//	fmt.Println(nums[i])
	//}

	x := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val: -1,
					Next: &ListNode{
						Val: -5,
						Next:nil,
					},
				},},
		},
	}
	y:=insertionSortList(&x)
	for y !=nil {
		fmt.Println(y.Val)
		y=y.Next
	}

}
