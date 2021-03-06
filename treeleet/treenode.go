package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"math"
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
	//root := TreeNode{data: 1}
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

	//root := TreeNode{data: 20}
	//root.left = NewNode(8);
	//root.left.left = NewNode(4);
	//root.left.right = NewNode(12);
	//root.left.right.left = NewNode(10);
	//root.left.right.right = NewNode(14);
	//root.right = NewNode(22);
	//root.right.right = NewNode(25);

	root := TreeNode{data: 8}
	root.left = NewNode(3)
	root.right = NewNode(10)
	root.left.left = NewNode(1)

	root.right.left = NewNode(6)
	root.right.right = NewNode(14)

	root.right.left.left = NewNode(4)
	root.right.left.right = NewNode(7)

	root.right.right.left = NewNode(13)

	return root
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

func printGivenLev(root *TreeNode, level int, ftr bool) {
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Printf("%d  ", root.data)
	} else if level > 1 {
		if ftr {
			printGivenLev(root.left, level-1, ftr)
			printGivenLev(root.right, level-1, ftr)

		} else {
			printGivenLev(root.right, level-1, ftr)
			printGivenLev(root.left, level-1, ftr)
		}
	}

}

func spiralPrint(root *TreeNode) {
	hight := height(root)
	ftr := true
	for i := 1; i <= hight; i++ {
		printGivenLev(root, i, ftr)
		ftr = !ftr
	}
}

// 变种二： Level order traversal with direction change after every two levels
/*
Given a binary tree, print the level order traversal in such a way that first two levels are printed from left to right, next two levels are printed from right to left, then next two from left to right and so on. So, the problem is to reverse the direction of level order traversal of binary tree after every two levels.

Examples:

Input:
            1
          /   \
        2       3
      /  \     /  \
     4    5    6    7
    / \  / \  / \  / \
   8  9 3   1 4  2 7  2
     /     / \    \
    16    17  18   19
Output:
1
2 3
7 6 5 4
2 7 2 4 1 3 9 8
16 17 18 19
In the above example, first two levels
are printed from left to right, next two
levels are printed from right to left,
and then last level is printed from
left to right.
 */
//方法一： 通过递归实现
func spiralPrint2(root *TreeNode) {
	hight := height(root)
	ftr := true
	cnt := 0
	for i := 1; i <= hight; i++ {
		printGivenLev(root, i, ftr)
		cnt ++
		if cnt == 2 {
			ftr = !ftr
			cnt = 0
		}
		fmt.Println()
	}
}

//方法二： 通过非递归实现 O(n)的时间复杂度
/*
1. 构建一个堆栈和一个queue，
2. 首先放入root到queue中，然后通过判断queue.len()>0: queu中的数据已经完全取完，即已经遍历完成
3. 然后从queue中取出数据，如果righttolef=false，则打印，如果true：则放入到stack中。其他的数据则放入queue中。
4. 然后判断stack的数据是否为空，如果不为空则取出然后打印，一直到stack为空
5. 每一层级的遍历都ct++ ,如果ct==2，则righttoleft= !righttoleft，然后ct=0
 */
func spiralPrint3(root *TreeNode) {
	q := queue.New()
	s := stack.New()
	qsize := 0
	ct := 0
	righttoleft := false
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		fmt.Printf("%d  ", root.data)
	}
	q.Enqueue(root)

	//首先放入root到queue中，然后通过判断queue.len()>0: queu中的数据已经完全取完，即已经遍历完成
	for q.Len() > 0 {
		ct ++
		//取出q的len，然后遍历queue，qsize代表这一层级的元素个数
		qsize = q.Len()
		//取出这一层级的元素
		for i := 0; i < qsize; i++ {
			qnode := q.Dequeue().(*TreeNode)
			//如果righttoleft为false，则打印，否则push到stack中
			if righttoleft {
				s.Push(qnode)
			} else {
				fmt.Printf("%d  ", qnode.data)
			}
			//然后放入下一级的数据到queue中。
			if qnode.left != nil {
				q.Enqueue(qnode.left)
			}
			if qnode.right != nil {
				q.Enqueue(qnode.right)
			}
		}

		//打印stack中的数据，达到反向打印的目的
		if righttoleft {
			for s.Len() > 0 {
				qnode := s.Pop().(*TreeNode)
				fmt.Printf("%d  ", qnode.data)
			}
		}
		if ct == 2 {
			righttoleft = !righttoleft
			ct = 0
		}
		fmt.Println()
	}

}

/*
func main() {
	root := InitTreeNode()
	spiralPrint2(&root)
	fmt.Println("使用queue和stack打印: \n\n")
	spiralPrint3(&root)

}
*/

//-------------------------------------------------------------------------------------------------------

/*
广度优先算法
 */
//----------------------------------------------------------------------------------------------
//计算tree的hight，有俩种方法：1，使用递归，2，不适用递归
/**

方法一： 使用递归计算tree的最大height：沿着根节点遍历到最长的叶子节点的长度
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

/**
不适用递归实现，借助queue来实现,
1. 放入root节点，nodecount ++
2. 循环条件为： nodecount >0 ;取出节点，如果存在left节点，入队，如果存在right节点，入队，然后nodecount--；
3. 外层循环为nodecount:=queue.Len()
4. 然后nodecount++
5. 然后继续取出，知道nodecount ==0时
 */
var levelCount int

func heightByQueue(root *TreeNode) int {
	levelCount = 0
	q := queue.New()
	//放入root
	q.Enqueue(root)
	for true {
		//nodecount 当前层的node数量
		nodecount := q.Len()
		if nodecount == 0 {
			return levelCount
		}
		//外层循环控制有多少层。当queue中没有节点时，代表所有节点都已经遍历了一遍。
		levelCount++
		//从queue中取出节点，当nodecount=0时，代表已经取出该层的节点了，
		for nodecount > 0 {
			//node := q.Peek()
			node := q.Dequeue().(*TreeNode)
			//如果node存在left节点，则入队，如果存在right节点则入队。
			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
			nodecount --
		}
	}
	return levelCount
}

//----------------------------------------------------------------------------------------------

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
		fmt.Printf("%d ", root.data)
	} else {
		printGivenLevel(root.left, level-1)
		printGivenLevel(root.right, level-1)
	}
}

//方法二非递归实现层级打印，借助queue来实现
func printLevelByQueue(root *TreeNode) {
	q := queue.New()
	q.Enqueue(root)
	for q.Len() > 0 {
		qnode := q.Dequeue().(*TreeNode)
		fmt.Printf("%d  ", qnode.data)
		if qnode.left != nil {
			q.Enqueue(qnode.left)
		}
		if qnode.right != nil {
			q.Enqueue(qnode.right)
		}
	}
}

/*
func main() {
	root := InitTreeNode()
	printLevelByQueue(&root)
	fmt.Println()
	printLevelWithoutResc(&root)
}
*/

//----------------------------------------------------------------------------------------------

/*
Print level order traversal line by line | Set 1

Example 2:
		 1
	  /     \
	 2       3
   /   \       \
  4     5       6
	   /  \     /
	  7    8   9
Output for above tree should be
1
2 3
4 5 6
7 8 9<

*/
// 方法一： 使用递归打印树，时间复杂度： O(n2)
func printLevelWithResc(root *TreeNode) {
	if root == nil {
		return
	}
	h := height(root)

	for i := 1; i <= h; i++ {
		printGivenLevel(root, i)
		fmt.Println()
	}
}

//方法二： 借助queue来实现: O(n)
func printLevelWithoutResc(root *TreeNode) {
	q := queue.New()
	q.Enqueue(root)
	level := 0
	for true {
		nodecount := q.Len()
		if nodecount == 0 {
			break
		}
		level ++

		for nodecount > 0 {
			node := q.Dequeue().(*TreeNode)
			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
			fmt.Printf("%d   ", node.data)
			nodecount --
		}
		fmt.Println()
	}
}

/*
func main() {
	root := InitTreeNode()
	fmt.Printf("递归height：   %d\n",height(&root))
	fmt.Printf("非递归height：   %d\n",heightByQueue(&root))

	root = TreeNode{}
	//-----------层级打印--------------------
	fmt.Println("非递归层级打印：   ")
	printLevelWithoutResc(&root)
	fmt.Println("\n\n")
	fmt.Println("非递归层级打印：   ")
	printLevelWithResc(&root)
}
*/

//---------------------------------------------------------------------------------------------
/*
反向打印level order travelal的数据
   					1
		2						3
	4			5				6		7
8       9   10     11       12     13

打印的结果：
8 9 10 11 12 13
4  5 6 7
2  3
1
*/
//使用传统方式打印
func reverseLevelOrder(root *TreeNode) {
	hgt := height(root)
	for i := hgt; i > 0; i-- {
		printGivenLev(root, i, true)
		fmt.Println()
	}
}

//采用stack方式
func reverseLevelOrderByStack(root *TreeNode) {
	q := queue.New()
	s := stack.New()
	if root == nil {
		return
	}
	if root.right == nil && root.left == nil {
		fmt.Printf("%d  ", root.data)
	}

	q.Enqueue(root)
	//s.Push(root)
	//for q.Len() >0 {
	//	qsize := q.Len()
	//	for i :=0;i<qsize;i++ {
	//		qnode := q.Dequeue().(*TreeNode)
	//
	//		if qnode.right !=nil {
	//			s.Push(qnode.right)
	//			q.Enqueue(qnode.right)
	//		}
	//		if qnode.left !=nil {
	//			s.Push(qnode.left)
	//			q.Enqueue(qnode.left)
	//		}
	//	}
	//	s.Push(nil)
	//}

	//通过queue，将所有节点反向压入到stack中，然后取出相关数据即可。
	for q.Len() > 0 {
		qnode := q.Dequeue().(*TreeNode)
		s.Push(qnode)
		if qnode.right != nil {
			q.Enqueue(qnode.right)
		}
		if qnode.left != nil {
			q.Enqueue(qnode.left)
		}
	}

	for s.Len() > 0 {
		tmp := s.Pop()
		if tmp != nil {
			qnode := tmp.(*TreeNode)
			fmt.Printf("%d   ", qnode.data)
		} else {
			fmt.Println()
		}
	}
}

//Perfect Binary Tree Specific Level Order Traversal
/*
https://www.geeksforgeeks.org/perfect-binary-tree-specific-level-order-traversal/
															1
								2															3
					4	    			 	5									6						7
			8				9		  10 			11					12				13		14				15
		16		17		18	   19  20 	  21	22		23			24		25		26		27 28	29		30		31


Print the level order of nodes in following specific manner:

  1 2 3 4 7 5 6 8 15 9 14 10 13 11 12 16 31 17 30 18 29 19 28 20 27 21 26  22 25 23 24
i.e. print nodes in level order but nodes should be from left and right side alternatively. Here 1st and 2nd levels are trivial.
While 3rd level: 4(left), 7(right), 5(left), 6(right) are printed.
While 4th level: 8(left), 15(right), 9(left), 14(right), .. are printed.
While 5th level: 16(left), 31(right), 17(left), 30(right), .. are printed.

一次处理俩个节点，都进行入队。
1. 入队： 第一个节点的左节点，第二个节点的右节点
2. 入队： 第二个节点的左节点，第一个节点的右节点

 */

func printSpecificLevelOrder(root *TreeNode) {
	q := queue.New()
	if root == nil {
		return
	}
	fmt.Printf("%d  ", root.data)
	if root.left != nil {
		fmt.Printf("%d   %d   ", root.left.data, root.right.data)
	}

	q.Enqueue(root.left)
	q.Enqueue(root.right)
	left := true
	for q.Len() > 0 {
		qnode1 := q.Dequeue().(*TreeNode)
		qnode2 := q.Dequeue().(*TreeNode)

		fmt.Printf("%d	%d	%d	%d	", qnode1.left.data, qnode2.right.data, qnode1.right.data, qnode2.left.data)
		if qnode1.left.left != nil {
			q.Enqueue(qnode1.left)
			q.Enqueue(qnode2.right)
			q.Enqueue(qnode1.right)
			q.Enqueue(qnode2.left)
			left = !left
		}
	}
}

/*
求tree的密度： 密度=size/height
 */
func heightAndSize(root *TreeNode) {
	q := queue.New()
	q.Enqueue(root)
	size := 0
	level := 0
	for q.Len() > 0 {
		qNum := q.Len()
		for i := 0; i < qNum; i++ {
			qnode := q.Dequeue().(*TreeNode)
			fmt.Printf("%d  ", qnode.data)
			size ++
			if qnode.left != nil {
				q.Enqueue(qnode.left)
			}
			if qnode.right != nil {
				q.Enqueue(qnode.right)
			}
		}
		fmt.Println()
		level++
	}
	fmt.Printf("size: %d ,   level:  %d,  密度： %f\n", size, level, math.Ceil(float64(size)/float64(level)))
}

//递归解法
func heightAndSizeRescUtil(root *TreeNode, size *int) int {
	if root == nil {
		return 0
	}
	llen := heightAndSizeRescUtil(root.left, size)
	rlen := heightAndSizeRescUtil(root.right, size)
	*size ++
	if llen > rlen {
		return llen + 1
	}
	return rlen + 1
}

func heightAndSizeResc(root *TreeNode) {
	var res *int
	var a int = 0
	res = &a
	height := heightAndSizeRescUtil(root, res)
	fmt.Println(math.Ceil(float64(*res) / float64(height)))
}

/*
LeetCode-53：Maximum Subarray (和最大的连续子数组)
 */
func maxSubAry(nums []int) int {
	futSum, endSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		endSum = max(endSum+nums[i], nums[i])
		futSum = max(futSum, endSum)
	}
	return futSum
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

/*
二叉树的边界遍历
Boundary Traversal of binary tree
1. 从上向下打印左子树
2. 打印叶子节点
3. 从下往上打印右子树
边界节点为： root.left !=nil时，root=root.left;否则为root=root.right
 */
func printBoundaryLeft(root *TreeNode) {
	if root != nil {
		if root.left != nil {
			fmt.Printf("%d  ", root.data)
			printBoundaryLeft(root.left)
		} else if root.right != nil {
			fmt.Printf("%d  ", root.data)
			printBoundaryLeft(root.right);
		}
	}
}

func printBoundaryRight(root *TreeNode) {
	if root != nil {
		if root.right != nil {
			printBoundaryRight(root.right)
			fmt.Printf("%d  ", root.data)
		} else if root.left != nil {
			printBoundaryRight(root.left)
			fmt.Printf("%d  ", root.data)
		}
	}
}

func printBoundaryLeaves(node *TreeNode) {
	if node != nil {
		printBoundaryLeaves(node.left)
		if node.left == nil && node.right == nil {
			fmt.Printf("%d  ", node.data)
		}
		printBoundaryLeaves(node.right)
	}

}

func printBoundary(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d  ", root.data)
		printBoundaryLeft(root.left)
		printBoundaryLeaves(root.left)
		printBoundaryLeaves(root.right)
		printBoundaryRight(root.right)
	}
}

/*
从preorder的数组计算出depth
Calculate depth of a full Binary tree from Preorder
Given preorder of a binary tree, calculate its depth(or height) [starting from depth 0]. The preorder is given as a string with two possible characters.

‘l’ denotes the leaf
‘n’ denotes internal node
The given tree can be seen as a full binary tree where every node has 0 or two children. The two children of a node can ‘n’ or ‘l’ or mix of both.
 */
func finddepth(root []rune) {
	var a int = 0
	var idx *int
	idx = &a
	n := len(root)
	fmt.Println(findDepthUtil(root, n, idx))
}
func findDepthUtil(root []rune, n int, idx *int) int {
	if root[*idx] == 'l' || *idx >= n {
		return 0
	}
	*idx ++
	llen := findDepthUtil(root, n, idx)
	*idx ++
	rllen := findDepthUtil(root, n, idx)
	if llen > rllen {
		return llen + 1
	}
	return rllen + 1
}

/*
Modify a binary tree to get preorder traversal using right pointers only
Given a binary tree. Modify it in such a way that after modification you can have a preorder traversal of it using only the right pointers. During modification, you can use right as well as left pointers.

Examples:

Input :    10
          /   \
        8      2
      /  \
    3     5
Output :    10
              \
               8
                \
                 3
                  \
                   5
                    \
                     2
Explanation : The preorder traversal
of given binary tree is 10 8 3 5 2.
 */

/*
Check for Children Sum Property in a Binary Tree
Given a binary tree, write a function that returns true if the tree satisfies below property.
For every node, data value must be equal to sum of data values in left and right children. Consider data value as 0 for NULL children. Below tree is an example
 */

func isSumProp(root *TreeNode) bool {
	q := queue.New()
	var left, right int
	q.Enqueue(root)
	for q.Len() > 0 {
		node := q.Dequeue().(*TreeNode)
		if node.left == nil {
			left = 0
		} else {
			left = node.left.data
			q.Enqueue(node.left)
		}
		if node.right == nil {
			right = 0
		} else {
			right = node.right.data
			q.Enqueue(node.right)
		}
		if node.left == nil && node.right == nil {
			continue
		} else {
			if node.data != (left + right) {
				return false
			}
		}
	}
	return true
}

/*
写一个function，用于判断是否是sumtree，sumtree:node = left node + right node
Write a function that returns true if the given Binary Tree is SumTree else false. A SumTree is a Binary Tree where the value of a node is equal to sum of the nodes present in its left subtree and right subtree. An empty tree is SumTree and sum of an empty tree can be considered as 0. A leaf node is also considered as SumTree.

Following is an example of SumTree.

          26
        /   \
      10     3
    /    \     \
  4      6      3

时间复杂度为： O(n2)
 */
func isSumTress(root *TreeNode) bool {

	if root == nil || root.left == nil && root.right == nil {
		return true
	}

	lsum := sum(root.left)
	rsum := sum(root.right)

	if (root.data == lsum+rsum) && isSumTress(root.left) && isSumTress(root.right) {
		return true
	}
	return false
}

func sum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return sum(node.left) + sum(node.right) + node.data
}

//方法二
/*
Method 2 ( Tricky )
The Method 1 uses sum() to get the sum of nodes in left and right subtrees. The method 2 uses following rules to get the sum directly.
1) If the node is a leaf node then sum of subtree rooted with this node is equal to value of this node.
如果是叶子节点： 叶子节点的和=node.data
2) If the node is not a leaf node then sum of subtree rooted with this node is twice the value of this node (Assuming that the tree rooted with this node is SumTree).
如果是非叶子节点: 节点的和 *2 =root.data

时间复杂度为O(n)
 */
func isSumTree(root *TreeNode) bool {
	if root == nil || root.left == nil && root.right == nil {
		return true
	}
	var lsum, rsum int
	for isSumTree(root.left) && isSumTree(root.right) {
		if root.left == nil {
			return true
		} else if root.left.left == nil && root.left.right == nil { //叶子节点
			lsum = root.left.data
		} else {
			lsum = 2 * root.left.data
		}

		if root.right == nil {
			return true
		} else if root.right.right == nil && root.right.left == nil {
			rsum = root.right.data
		} else {
			rsum = 2 * root.right.data
		}
		return root.data == lsum+rsum
	}
	return false
}

/*
三个数中有一个是负数的时候就需要小心：需要比对三个最大数和前俩个最小数*最大数的大小。
 */
func maxSum(arr []int, n int) int {
	if len(arr) < n {
		return 0
	}
	sum := 1
	j := 0
	for j < n {
		max := arr[0]
		i := 0
		for ; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		//fmt.Println("i: ",i)
		if i == len(arr) {
			arr = arr[0 : i-1]
		} else {
			arr = append(arr[0:i-2], arr[i-1:len(arr)-1]...)
		}
		sum *= max
		j++
	}
	return sum
}

/*
把二叉树按照对角线打印
d： 代表坡度
使用map来保存每个坡度对应的值

需要按照preorder来打印，先打印root
 */
func diagonalPrintUtil(root *TreeNode, d int, m map[int][]int) {
	if root == nil {
		return
	}
	tmp := append(m[d], root.data)
	m[d] = tmp
	diagonalPrintUtil(root.left, d+1, m)
	diagonalPrintUtil(root.right, d, m)
}

func diagonalPrint(root *TreeNode) {
	m := make(map[int][]int)
	diagonalPrintUtil(root, 0, m)

}

// 迭代实现：
/*
使用preorder：
1.打印root和right节点
2. 压入left节点
3.如果right节点未nil的话，继续弹出节点，判断是否为空，为空则换行打印，标志位一个对角线结束，然后继续压入nil
 */
func diagonalPrintWithoutResc(root *TreeNode) {
	q := queue.New()
	if root == nil {
		return
	}
	q.Enqueue(root)
	q.Enqueue(nil)

	for q.Len() > 0 {
		p := q.Dequeue()
		var tmp *TreeNode
		if p != nil {
			tmp = p.(*TreeNode)
		} else {
			tmp = nil
		}
		if tmp == nil {
			if q.Len() == 0 {
				return
			}
			fmt.Println()
			q.Enqueue(nil)
		} else {
			for tmp != nil {
				fmt.Printf("%d ", tmp.data)
				if tmp.left != nil {
					q.Enqueue(tmp.left)
				}
				tmp = tmp.right
			}
		}
	}

}




/*

94. 二叉树的中序遍历
给定一个二叉树，返回它的中序 遍历。


 */
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	helper(root, &ret)
	return ret
}

func helper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	helper(root.left, ret)
	*ret = append(*ret, root.data)
	helper(root.right, ret)
}


/*
96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func numTrees(n int) int {
	//dp[i]=dp[j]+dp[j-1]
	dp := make([]int,n+1)
	dp[0]=1
	dp[1]=1

	for i:=2;i<n+1;i++ {
		for j:=1;j<i+1;j++ {
			dp[i]+=dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

/*
95. 不同的二叉搜索树 II
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:

输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//func generateTrees(n int) []*TreeNode {
//
//	return genTree(1,n)
//}
//
//func genTree(start int,end int) []*TreeNode {
//	ret := []*TreeNode{}
//	if start > end {
//		return ret
//	}
//	if start == end {
//		node=&TreeNode{Val:start}
//		ret=append(ret,node)
//		return node
//	}
//	var lnode ,rnode *TreeNode
//	for i:=start;i<=end;i++ {
//		lnode = genTree(i-1,end)
//		rnode=genTree(i+1,end)
//		for _,ln := range lnode{
//			for _,rn :=range rnode {
//				root := &TreeNode{Val:i}
//				root.left=ln
//				root.right =rn
//				ret =append(root)
//			}
//		}
//	}
//}


func main() {

	//root := InitTreeNode()
	//fmt.Println(inorderTraversal(&root))
	fmt.Println(numTrees(3))
	//reverseLevelOrder(&root)
	//reverseLevelOrderByStack(&root)
	//printSpecificLevelOrder(&root)
	//fmt.Println()
	//heightAndSize(&root)
	//heightAndSizeResc(&root)
	//
	//nums := []int{-1, 1, 2, -1, 1, 3, -1, 3}
	//fmt.Println(maxSubAry(nums))
	//printBoundary(&root)
	//fmt.Println()
	//var runes []rune = []rune{'n', 'l', 'n', 'n', 'l', 'l', 'l'}
	//finddepth(runes)

	//fmt.Println(isSumProp(&root))
	//fmt.Println(isSumTree(&root))
	//
	//arr := []int{-6, -5, -1, -2, -3, -4, 5, 6}
	//fmt.Println(maxSum(arr, 3))

	//diagonalPrintWithoutResc(&root)
}
