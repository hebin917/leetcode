package main

import "fmt"

type node struct {
	data  rune
	left  *node
	right *node
}

func newnode(data rune) (n *node) {
	n = &node{
		data: data,
	}
	n.left = nil
	n.right = nil
	return n
}

func initNode() (*node) {
	/**
				1
		2				3
	4       5       6        7
8
	 */
	root := newnode('A')
	//root.left = NewNode(2)
	//root.right = NewNode(3)
	//
	//root.left.left = NewNode(4)
	//root.left.right = NewNode(5)
	//
	//root.right.left = NewNode(6)
	//root.right.right = NewNode(7)
	//
	//root.left.left.left = NewNode(8)
	//root.left.left.right = NewNode(9)
	//
	//root.left.right.left = NewNode(10)
	//root.left.right.right = NewNode(11)
	//
	//root.right.left.left = NewNode(12)
	//root.right.left.right = NewNode(13)
	//
	//root.right.right.left = NewNode(14)
	//root.right.right.right = NewNode(15)

	return root
}

/*
Construct Tree from given Inorder and Preorder traversals
Let us consider the below traversals:

Inorder sequence: D B E A F C
Preorder sequence: A B D E C F

In a Preorder sequence, leftmost element is the root of the tree. So we know ‘A’ is root for given sequences. By searching ‘A’ in Inorder sequence, we can find out all elements on left side of ‘A’ are in left subtree and elements on right are in right subtree. So we know below structure now.

                 A
               /   \
             /       \
           D B E     F C
We recursively follow above steps and get the following tree.

         A
       /   \
     /       \
    B         C
   / \        /
 /     \    /
D       E  F

 */

var (
	preIndex int
)

func constructTree(in []rune, pre []rune, start, end int) *node {
	if start > end  || preIndex>end {
		return nil
	}
	fmt.Printf("preindex: %d   predata: %q   st:  %d   end:  %d \n",preIndex,pre[preIndex],start,end)
	//rootdata :=pre[preIndex]end:  "
	rootdata := pre[preIndex]
	rootnode := newnode(rootdata)
	preIndex++
	if start == end {
		return rootnode
	}

	idx := search1(in,rootdata,start,end)
	rootnode.left=constructTree(in,pre,start,idx-1)
	rootnode.right=constructTree(in,pre,idx+1,end)


	return rootnode
}

func search1(in []rune, key rune, start, end int) int {
	for i := start; i < end; i++ {
		if in[i] == key {
			return i
		}
	}
	return -1
}

func printInorderTree(root *node) {
	if (root == nil) {
		return
	}

	/* first recur on left child */
	printInorderTree(root.left)

	/* then print the data of node */
	fmt.Printf("%q  ", root.data)

	/* now recur on right child */
	printInorderTree(root.right)
}

func main() {
	in := []rune{'D', 'B', 'E', 'A', 'F', 'C'}
	pre := []rune{'A', 'B', 'D', 'E', 'C', 'F'}
	preIndex =0
	root := constructTree(in, pre, 0, len(pre) - 1)

	printInorderTree(root)
}
