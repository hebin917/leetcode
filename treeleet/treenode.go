package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	)

/*
二叉树特性：
1. 每层最多有2的l次方-1 个节点
2. 最多有： 2的l -1 个节点
3. 最小的高度为： log2（n+1）
4. 有L个叶子节点，那么最小的层级： log2L +1
5. 如果binary每个节点只有2或0个child，那么叶子节点的个数为nodes的个数+1：  L=N+1

二叉树种类
1. Full Binary完全二叉树，完整的二叉树为一个二叉堆
2. Perfect Binary Tree：所有内部节点都被填满，且所有的叶子都在相同的层级
3. 平衡二叉树： 树的高度为O(log n ),n 为nodes的数量，则该树为平衡二叉树
AVL tree maintains O(Log n) height by making sure that the difference between heights of left and right subtrees is 1. Red-Black trees maintain O(Log n) height by making sure that the number of Black nodes on every root to leaf paths are same and there are no adjacent red nodes. Balanced Binary Search trees are performance wise good as they provide O(log n) time for search, insert and delete.

4. 退化树： 每个树节点只有一个子节点，与链表相同

 */
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func NewNode(data int) (n *TreeNode) {
	n = &TreeNode{
		data: data,
	}
	n.left = nil
	n.right = nil
	return n
}

func InitTreeNode() (TreeNode) {
	/**
				1
		2				3
	4       5       6        7
8
	 */
	root := TreeNode{data: 1}
	root.left = NewNode(2)
	root.right = NewNode(3)

	root.left.left = NewNode(4)
	root.left.right = NewNode(5)

	root.right.left = NewNode(6)
	root.right.right = NewNode(7)

	root.left.left.left = NewNode(8)
	return root
}

//----------------------------------------------------------------------------------------------

/**
计算tree的最大height：沿着根节点遍历到最长的叶子节点的长度
 */
func height(root *TreeNode) int {
	if root == nil {
		return 0;
	} else {
		llen := height(root.left)
		rlen := height(root.right)
		if llen < rlen {
			return rlen + 1
		} else {
			return llen + 1
		}
	}
}

/*
按照层级打印： breadth first traversal： 即广度优先算法
1. 先找出tree的深度
2. 然后根据层级来进行打印
3. 层级打印： printGivenLevel ： 递归基：level=1,然后打印root.data
 */
func printLevel(root *TreeNode) {
	if root == nil {
		return
	}
	len := height(root)

	for i := 1; i <= len; i++ {
		printGivenLevel(root, i)
	}
}

func printGivenLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == 1 {
		print("%d ", root.data)
	} else {
		printGivenLevel(root.left, level-1)
		printGivenLevel(root.right, level-1)
	}
}

//----------------------------------------------------------------------------------------------

/**
树遍历的方式：
1. breath first traversal： 广度优先
2. depth first traversal： 深度优先

遍历: 以root节点打印的顺序分成三种： 前中后序遍历，利用递归来做
1. Inorder： Left ROOt，Right来打印： 打印出的顺序是左/根/右 ：用于二叉搜索树，  Inorder遍历以非递减顺序给出节点。为了以非递增的顺序获得BST的节点，可以使用Inorder遍历反转的Inorder遍历的变体
2. PreOder: 前序遍历： 根、左节点、右节点      ： 用于创建树的一个copy，也可以被用于获取表达式树的前缀表达式
3. PostOrder： 后续遍历： 左节点/右节点/Root  : 用于删除树

复杂度： T(n)=T(k) + T(n-k-1) +c ，其中k:在root一端的k个节点，那么另外一端为： n-k-1 。 c: root

Let’s do an analysis of boundary conditions

Case 1: Skewed tree (One of the subtrees is empty and other subtree is non-empty )

k is 0 in this case.
T(n) = T(0) + T(n-1) + c
T(n) = 2T(0) + T(n-2) + 2c
T(n) = 3T(0) + T(n-3) + 3c
T(n) = 4T(0) + T(n-4) + 4c

…………………………………………
………………………………………….
T(n) = (n-1)T(0) + T(1) + (n-1)c
T(n) = nT(0) + (n)c

Value of T(0) will be some constant say d. (traversing a empty tree will take some constants time)

T(n) = n(c+d)
T(n) = Θ(n) (Theta of n)

Case 2: Both left and right subtrees have equal number of nodes.

T(n) = 2T(|_n/2_|) + c

This recursive function is in the standard form (T(n) = aT(n/b) + (-)(n) ) for master method http://en.wikipedia.org/wiki/Master_theorem. If we solve it by master method we get (-)(n)

Auxiliary Space : If we don’t consider size of stack for function calls then O(1) otherwise O(n).

 */

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d  ", root.data)
	inorder(root.right)
}

/*
1. 非递归遍历二叉树，使用stack来保存最左边的节点: cur = cur.left
2. 从stack中弹出相关节点，然后打印，这样就能打印出左/根，然后就缺右节点
3. 当stack不为空时，放入cur=cur.right.这样又会回到1的逻辑，一直放入left节点，然后又从后往前取出左/根/右
4. 当stack为空的时候，表明整个节点已经访问完成，设置done=false结束循环。
 */
func inorderWithoutRecursion(root *TreeNode) {
	s := stack.New()
	cur := root
	done := true
	for ; done; {
		//1. 先将左边的节点全部入队
		if cur != nil {
			s.Push(cur)
			cur = cur.left
		} else {
			//2. 如果stack不为NULL，则取出数据，打印，然后将cur=cur.right，然后回到1，如果cur不为空，则放入右节点的left节点，然后一直循环1和2，直到所有的节点都被放入和取出
			if s.Len() != 0 {
				tmp := s.Pop().(*TreeNode)
				fmt.Printf("%d  ", tmp.data)
				cur = tmp.right
			} else {
				done = false
			}
		}

	}
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d  ", root.data)
	preOrder(root.left)
	preOrder(root.right)
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d  ", root.data)
}

/*
func main() {
	root := InitTreeNode()
	inorder(&root)
	fmt.Println()
	fmt.Println(height(&root))

	printLevel(&root)

	fmt.Printf("前序遍历: \n")
	preOrder(&root)

	fmt.Printf("中序遍历: \n")
	inorder(&root)

	fmt.Printf("后序遍历: \n")
	postOrder(&root)

	fmt.Println()
	fmt.Println("中序遍历：  非递归")
	inorderWithoutRecursion(&root)
}
*/

//----------------------------------------------------------------------------------------------

/*
Print Postorder traversal from given Inorder and Preorder traversals
给定inorder和preorder的数组，然后打印出postorder的数据

思路：
1. preorder中，第一个元素为root
2. Inorder以root相隔，找到root的index，左边的未左节点，右边的为右节点
3. 然后查询preorder中的下一个元素，他可能是叶子节点的父节点，然后依次找到左右子树
 */

// 查找父节点的方法: n 为要遍历的长度
func search(arr [] int, x int, len int) int {
	for i := 0; i < len; i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}

func printPostOrder(in [] int, pre [] int, len int) {
	root := search(in, pre[0], len)

	//root ！=0 代表左节点不为空，则打印左数
	if root != 0 {
		printPostOrder(in[0:root], pre[1:len], root)
	}

	//如果右子树不为空，则打印右子树
	if root != len-1 {
		printPostOrder(in[root+1:len], pre[root+1:], len-root-1)
	}
	//打印root
	fmt.Printf("%d  ", pre[0])
}

//方法二：

func search2(in []int, startIn, endIn, data int) int {
	i := 0
	for i = startIn; i < endIn; i ++ {
		if in[i] == data {
			return i
		}
	}
	return i
}

var preIndex int = 0

func printPostOrder2(in [] int, pre [] int, inStr, inEnd int) {
	if inStr > inEnd {
		return
	}

	inIndex := search2(in, inStr, inEnd, pre[preIndex])
	preIndex ++
	//则打印左数,左子树为inStr到父节点的位置
	printPostOrder2(in, pre, inStr, inIndex-1)

	//则打印右子树

	printPostOrder2(in, pre, inIndex+1, inEnd)

	//打印root
	fmt.Printf("%d  ", in[inIndex])
}

/*
func main() {
	pre := []int{1, 2, 4, 5, 3, 6}
	in := []int{4, 2, 5, 1, 3, 6}
	//root := pre[0]
	len := len(pre)
	fmt.Println("方法一： ")
	printPostOrder(in, pre, len)
	fmt.Println()
	fmt.Println("方法二： ")
	printPostOrder2(in, pre, 0, len-1)
}
*/

//-------------------------------------------------------------------------------------------------------
/**
Find postorder traversal of BST from preorder traversal
Given an array representing preorder traversal of BST, print its postorder traversal.
从二叉排序树的preorder中找出postorder，给定一个preorder的array，然后打印出postorder的traversal

BST： Binary Search Tree，称为二叉排序树，或者二叉查找树，树的值在分布的时候具有非常明显的特征，左子树的值小于根节点的值，而根节点的值小于右子树的值，二叉排序树本身是具有动态性的，可以动态地进行节点的删除，插入等的操作

时间复杂度： O(h),h为树的高度。


Examples:

Input : 40 30 35 80 100
Output : 35 30 100 80 40

Input : 40 30 32 35 80 90 100 120
Output : 35 32 30 120 100 90 80 40

方法： 找到根节点，root=40，然后找出比根大的值的index，index的左边为左子树，index右边的为右子树，依次根据父节点来进行查找
 */

func searchPre(pre []int, strt int, end int, parentData int) int {
	for i := strt; i < end; i++ {
		if pre[i] > parentData {
			return i
		}
	}
	return -1
}

func constructPost(pre []int, size int) *TreeNode {
	preIndex := 0
	return constructPostUtil(pre, preIndex, 0, size-1, size)
}

// A recursive function to construct Full from pre[]. preIndex is used
// to keep track of index in pre[].
func constructPostUtil(pre []int, preIndex1 int, low int, high int, size int) *TreeNode {
	if preIndex >= size && low >= high {
		return nil
	}

	root := &TreeNode{data: pre[preIndex1]}
	preIndex = preIndex1 + 1

	//If the current subarry has only one element, no need to recur
	if low == high {
		return root
	}

	//求出左右节点的分界点
	idx := low + 1
	for ; idx <= high; idx++ {
		if pre[idx] > root.data {
			break
		}
	}
	//左节点为： root到index-1 的节点部分
	root.left = constructPostUtil(pre, preIndex, preIndex, idx-1, size)

	//右节点： index 至结束的位置
	root.right = constructPostUtil(pre, preIndex, idx, high, size)

	return root;
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.left)
	fmt.Printf("%d   ", root.data)
	printInorder(root.right)
}

/*
func main() {
	pre := []int{10, 5, 1, 7, 40, 50}
	len := len(pre)
	root := constructPost(pre, len)
	printInorder(root);
}
*/
//-------------------------------------------------------------------------------------------------------
/*
Find n-th node of inorder traversal
Given the binary tree and you have to find out the n-th node of inorder traversal.
 */
var count int = 0

func NthInorder(root *TreeNode, nth int) {
	if root == nil {
		return
	}
	if count <= nth {
		NthInorder(root.left, nth)
		count++
		if count == nth {
			fmt.Printf("%d-th node is:   %d", nth, root.data)
		}
		NthInorder(root.right, nth)
	}
}

var postCount int = 0

func NthPostOrder(root *TreeNode, nth int) {
	if root == nil {
		return
	}
	if postCount <= nth {
		NthPostOrder(root.left, nth)
		NthPostOrder(root.right, nth)
		postCount ++
		if postCount == nth {
			fmt.Printf("\n%d-th postorder is: %d", nth, root.data)
		}
	}
}

/*
func main() {
	root := InitTreeNode()
	NthInorder(&root, 4)
	NthPostOrder(&root,4)
}*/

//-------------------------------------------------------------------------------------------------------
/*
Level order traversal in spiral form
Write a function to print spiral order traversal of a tree. For below tree, function should print 1, 2, 3, 4, 5, 6, 7.
写一个函数以螺旋形打印tree，每一层级的打印方向不一样
 */

func printlevel(root *TreeNode, level int) {

}

func printGivenLev(root *TreeNode, level int,ftr bool) {
	if root ==nil {
		return
	}
	if level == 1 {
		fmt.Printf("%d  ",root.data)
	}else if level >1 {
		if ftr {
			printGivenLev(root.left,level-1,ftr)
			printGivenLev(root.right,level-1,ftr)

		}else {
			printGivenLev(root.right,level-1,ftr)
			printGivenLev(root.left,level-1,ftr)
		}
	}

}

func spiralPrint(root *TreeNode) {
	hight := height(root)
	ftr :=true
	for i:=1;i<=hight;i++ {
		printGivenLev(root,i,ftr)
		ftr = !ftr
	}
}

/*
func main() {
	root := InitTreeNode()
	spiralPrint(&root)

}
*/

//-------------------------------------------------------------------------------------------------------

/**
插入一个节点，在为NULL的时候
*/
func insert(node *TreeNode, key int) {

}
