package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	root := InitTreeNode()
	inorder(&root)
	fmt.Println()
	IterInorder(&root)


	fmt.Println("PreOder:  ")
	preOrder(&root)
	fmt.Println("")
	IterPreorder(&root)

	fmt.Println()
	fmt.Println("PostOder:  ")
	postOrder(&root)

	fmt.Println("PostOrder iter: ")
	IterPostorderWithTwoStack(&root)
	heightAndSize(&root)
}



/*
使用stack来操作
1. 当cur不为空时，压入cur，然后cur=cur.left
2. 当cur为空时，代表左子树已经压完，
3。stack.pop，打印，然后置换cur=cur.right，开始压右子树。
4. 重复1，如果右子树不为空，则压入右子树的左子树，一直重复1，2，3步

 */
func IterInorder(root *TreeNode) {
	s := stack.New()
	cur := root
	for (cur != nil || s.Len() > 0) {
		for cur != nil {
			s.Push(cur)
			cur = cur.left
		}
		cur = s.Pop().(*TreeNode)
		fmt.Printf("%d  ", cur.data)
		cur = cur.right
	}
	fmt.Println()

}

/*
1. 压入root
2. 弹出cur:= s.Pop().(*TreeNode)，并打印
3. 判断cur.right是否为空，如果不为空，则压入cur.right
4. 判断cur.left是否为空，如果不为空，则压入cur.left

5. 重复第一步，则先压入左栈所有元素，弹出打印，然后取出右栈，弹出打印
 */
func IterPreorder(root *TreeNode) {
	s:=stack.New()
	s.Push(root)
	for s.Len()>0 {
		cur := s.Pop().(*TreeNode)
		fmt.Printf("%d  ",cur.data)
		if cur.right !=nil{
			s.Push(cur.right)
		}
		if (cur.left !=nil ){
			s.Push(cur.left)
		}
	}
}

/*

 */
func IterPostorderWithOneStack(root *TreeNode) {
	if root ==nil {
		return
	}
	s :=stack.New()
	for s.Len() >0 {
		for root !=nil{
			if root.right !=nil {
				s.Push(root.right)
			}
			s.Push(root)
			root=root.left
		}

		root := s.Pop().(*TreeNode)
		if root.right !=nil && s.Peek().(*TreeNode) == root.right {
			s.Pop()
			s.Push(root)
			root=root.right
		}else {
			fmt.Printf("%d  ",root.data)
			root=nil
		}
	}
}


/*

 */
func IterPostorderWithTwoStack(root *TreeNode) {
	if root ==nil {
		return
	}
	s1 := stack.New()
	s2 := stack.New()

	s1.Push(root)
	for s1.Len() >0 {
		cur := s1.Pop().(*TreeNode)
		s2.Push(cur)

		if cur.left !=nil {
			s1.Push(cur.left)
		}
		if cur.right != nil {
			s1.Push(cur.right)
		}
	}

	for s2.Len() >0 {
		fmt.Printf("%d  ",s2.Pop().(*TreeNode).data)
	}

}

/*
Diagonal Traversal of Binary Tree:  二叉树的对角遍历

 */