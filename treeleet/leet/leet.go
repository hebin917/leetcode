package main

import (
	"math"
	"github.com/eapache/queue"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"github.com/golang-collections/collections/set"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	ret := []*TreeNode{}
	genTree(1, n, ret)
	return ret
}

func genTree(start int, end int, rets []*TreeNode) []*TreeNode {
	ret := []*TreeNode{}
	if start > end {
		return ret
	}
	if start == end {
		node := &TreeNode{Val: start}
		ret = append(ret, node)
		return ret
	}

	var lnode, rnode []*TreeNode
	for i := start; i <= end; i++ {
		lnode = genTree(start, i-1, rets)
		rnode = genTree(i+1, end, rets)
		for _, ln := range lnode {
			for _, rn := range rnode {
				root := &TreeNode{Val: i}
				root.Left = ln
				root.Right = rn
				rets = append(rets, root)
			}
		}
	}
	return ret
}

/*
98. 验证二叉搜索树

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
 */
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt32, math.MaxInt32)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

/*
102. 二叉树的层次遍历

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	q := queue.New()
	q.Add(root)
	for q.Length() > 0 {
		qlen := q.Length()
		tmp := make([]int, qlen)
		for i := 0; i < qlen; i++ {

			cur := q.Remove().(*TreeNode)
			tmp[i] = cur.Val
			if cur.Left != nil {
				q.Add(cur.Left)
			}
			if cur.Right != nil {
				q.Add(cur.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

/*
103. 二叉树的锯齿形层次遍历
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var (
		res  [][]int = [][]int{}
		flag bool    = true
	)
	q := queue.New()

	q.Add(root)
	for q.Length() > 0 {
		//初始化每层的tmp数量
		qlen := q.Length()
		tmp := make([]int, qlen)
		for i := 0; i < qlen; i++ {
			node := q.Remove().(*TreeNode)
			index := i
			//如果flag为true，则创建
			if !flag {
				index = qlen - i - 1
			}

			tmp[index] = node.Val
			fmt.Printf("%d  ", node.Val)

			if node.Left != nil {
				q.Add(node.Left)
			}
			if node.Right != nil {
				q.Add(node.Right)
			}

		}
		fmt.Println()
		flag = !flag
		res = append(res, tmp)
	}
	return res
}

/*
105. 从前序与中序遍历序列构造二叉树

根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	//前序是root ，左 右
	// 中序： 左，右  中
	// 可以通过中序拿到par，然后在inorder中找到par所在的index，然后根据index可以区分出左和右，然后继续遍历
	var root *TreeNode = &TreeNode{Val: preorder[0]}
	if len(inorder) == 0 {
		return nil
	}
	//求出parent节点所在的index
	index := searchTree(inorder, preorder[0])

	//index左边的未左节点
	root.Left = buildTree(inorder[:index], preorder[1:index+1])

	//index右边的为右节点
	root.Right = buildTree(inorder[index+1:], preorder[index+1:])
	return root
}

func searchTree(inorder []int, val int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			return i
		}
	}
	return 0
}

/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func consuctTree(inorder []int, postorder []int) *TreeNode {
	//中序可以区分左右节点
	//后续： 左，右 根：  根为最后一个节点，
	//左节点: inorder: :index     postorder: :index
	//🈶节点： inorder:  index+1:   postorder:  index+1
	//postorder中需要移除root节点\
	if len(inorder) == 0 {
		return nil
	}
	plen := len(postorder)
	root := &TreeNode{Val: postorder[plen-1]}
	index := searchTree(inorder, postorder[plen-1])

	root.Left = consuctTree(inorder[:index], postorder[:index])
	root.Right = consuctTree(inorder[index+1:], postorder[index:plen-1])
	return root
}

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	//初始化一个min_dept为最大值，如果left<min_dept，那么设置min_dept=left。然后如果right < min_dept，则设置min_dept=right
	min_dept := math.MaxInt32
	if root.Left != nil {
		left := minDepth(root.Left)
		if left < min_dept {
			min_dept = left
		}
	}

	if root.Right != nil {
		right := minDepth(root.Right)
		if right < min_dept {
			min_dept = right
		}
	}

	return min_dept + 1
}

/*
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int = [][]int{}
	var helper func(root *TreeNode, sum int, path []int)
	helper = func(root *TreeNode, sum int, path []int) {
		if root == nil {
			return
		}
		sum -= root.Val
		path = append(path, root.Val)
		//结果就是在叶子节点且sum=0时，才将tmp数组的值放入到res中。

		if root.Left == nil && root.Right == nil && sum == 0 {
			res = append(res, append([]int(nil), path...))
		}
		helper(root.Right, sum, path)
		helper(root.Left, sum, path)
	}
	//slice是引用传值，而数组是值传递
	helper(root, sum, []int{})
	return res
}

/*
110 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dept(root) != -1
}

func dept(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dept(root.Left)
	if left == -1 {
		return -1
	}
	right := dept(root.Right)
	if right == -1 {
		return -1
	}
	x := float64(left - right)
	if math.Abs(x) < 2 {
		if left > right {
			return left + 1
		}
		return right + 1
	}
	return -1
}

/*
112. 路径总和
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例: 
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum = sum - root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

/*
114将tree转化为链表
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func flatten(root *TreeNode) {
	//思路：将链表展开，反转后，就是树的postorder的顺序,并且逆序给出，这样的话应该先访问right，然后访问left，最后访问root，所以应用postorder方式去访问树
	//先访问right，然后访问left;通过prev来保存已经访问的元素
	//然后root.right=prev,root.left=nil。prev=root
	if root == nil {
		return
	}
	var prev *TreeNode
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Right)
		helper(root.Left)
		root.Right = prev
		root.Left = nil
		prev = root
	}
	helper(root)
}

/*
116. 填充每个节点的下一个右侧节点指针
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type Node struct {
	val   int
	left  *Node
	right *Node
	next  *Node
}

func connect(root *Node) {
	if root == nil {
		return
	}
	q := stack.New()
	q.Push(root)
	var node *Node
	for q.Len() > 0 {
		qlen := q.Len()
		var prev *Node = nil
		for i := qlen; i > 0; i-- {
			node = q.Pop().(*Node)
			node.next = prev
			prev = node
			if node.right != nil {
				q.Push(node.left)
			}
			if node.left != nil {
				q.Push(node.right)
			}
		}
	}
}

//相当于只用访问node.left即可，每次访问node.left将node.left指向node.right
func connect1(root *Node) {

}

/*
129. 求根到叶子节点数字之和
给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。

例如，从根到叶子节点路径 1->2->3 代表数字 123。

计算从根到叶子节点生成的所有数字之和。

说明: 叶子节点是指没有子节点的节点。

示例 1:

输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.
示例 2:

输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-root-to-leaf-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func sumNumbers(root *TreeNode) int {
	//采用俩个stack，s1: 压入node相关的元素，s2： 压入path
	//在叶子节点的时候进行相加
	if root == nil {
		return 0
	}
	nodest, pathst := stack.New(), stack.New()
	var sum int64
	nodest.Push(root)
	pathst.Push(fmt.Sprintf("%d", root.Val))
	for nodest.Len() > 0 {
		node := nodest.Pop().(*TreeNode)
		path := pathst.Pop().(string)

		if node.Left != nil {
			nodest.Push(node.Left)
			pathst.Push(fmt.Sprintf("%s%d", path, node.Left.Val))
		}
		if node.Right != nil {
			nodest.Push(node.Right)
			pathst.Push(fmt.Sprintf("%s%d", node.Right.Val))
		}

		if node.Left == nil && node.Right == nil {
			n, _ := strconv.ParseInt(path, 10, 32)
			sum += n
		}

	}
	return int(sum)
}

func sumNumbersWithResc(root *TreeNode) int {
	var (
		helper func(root *TreeNode, sum int) int
	)
	helper = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			//代表叶子节点
			return sum
		}
		left := helper(root.Left, sum)
		right := helper(root.Right, sum)
		return left + right
	}
	return helper(root, 0)
}

/*
404. 左叶子之和
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */
var sum int = 0

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sumLeftLeaves(root, sum)
	return sum
}
func sumLeftLeaves(root *TreeNode, sum int) {
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		} else {
			sumLeftLeaves(root.Left, sum)
		}
	}
	if root.Right != nil {
		sumLeftLeaves(root.Right, sum)
	}
}

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	h := height(root)
	fmt.Println("heigt:  ", h)
	res := [][]int{}
	for i := h; i >= 1; i-- {
		tmp := []int{}
		getGivenLevel(root, i, &tmp)
		res = append(res, tmp)
	}

	return res
}

func getGivenLevel(root *TreeNode, level int, levRes *[]int) {
	if root == nil {
		return
	}
	if level == 1 {
		*levRes = append(*levRes, root.Val)
	} else if level > 1 {
		getGivenLevel(root.Left, level-1, levRes)
		getGivenLevel(root.Right, level-1, levRes)
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

/*
236 二叉树的最低公共节点
1. 使用stack遍历整棵树
2. 将节点的parent放入map中
3. 然后遍历map，将p所在的parent放入一个set中
4. 遍历另外一个节点，如果在parent所在的set中，然后就返回即可。
 */
func lowestCommonAnnosector(root *TreeNode, p, q *TreeNode) *TreeNode {
	var (
		parm   map[*TreeNode]*TreeNode = make(map[*TreeNode]*TreeNode)
		annset set.Set                 = set.Set{}
	)
	st := stack.New()

	st.Push(root)
	parm[root] = nil
	_, ok1 := parm[p]
	_, ok2 := parm[q]
	for !ok1 || ! ok2 {
		_, ok1 = parm[p]
		_, ok2 = parm[q]
		node := st.Pop().(*TreeNode)
		if node.Left != nil {
			st.Push(node.Left)
			parm[node.Left] = node
		}

		if node.Right != nil {
			st.Push(node.Right)
			parm[node.Right] = node
		}
	}
	//如果p不等于nil，然后将其放入annset，然后一直将p的父节点放入到annset
	for p != nil {
		annset.Insert(p)
		p = parm[p]
	}

	//判断annset中是否存在q，如果不存在一直取出q的父节点来访问，如果q的父节点存在于annset中的话，就表明他们存在公共节点，且公共节点未q或者q的父节点。
	for !annset.Has(q) {
		q = parm[q]
	}
	return q
}

/*
235 二叉搜索树的最低公共子节点
解法：因为是二叉搜索树，左节点会小于根及诶单，然后右节点会大于根节点。如果一个在左,一个在右，那么最低公共子节点未root。如果都在一边，则判断和root值的大小
left.val < root.val :  root.val -left.val >0
right.val < root.val:  root.val -right.val <0
如果都在一边，则俩个的乘积应该大于0
 */
func lowestCommonAnnsector(root *TreeNode, left, right *TreeNode) *TreeNode {
	for (root.Val-left.Val)*(root.Val-right.Val) > 0 {
		if left.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

/*
124. 二叉树中的最大路径和
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//设置sum表示最后的返回值
	sum := math.MinInt32
	var helper func(root *TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		//获取左右节点的maxpathsum值，
		left := max(helper(root.Left), 0)
		right := max(helper(root.Right), 0)

		// the price to start a new path where `node` is a highest node
		price_newpath := root.Val + left + right;

		// update max_sum if it's better to start a new path
		sum = max(sum, price_newpath)

		// for recursion :
		// return the max gain if continue the same path
		return root.Val + max(left, right)
	}
	helper(root)
	return sum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/*
173. 二叉搜索树迭代器
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 next() 将返回二叉搜索树中的下一个最小的数。
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false
 

提示：

next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。

//解法：
1. 通过stack来实现，
2.构造时，先压入左边的数
3， 每次next的时候返回pop，弹出后压入右边的数
 */
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	st := []*TreeNode{}
	cur := root
	st = append(st, cur)
	for cur != nil {
		st = append(st, cur)
		cur = cur.Left
	}
	return BSTIterator{
		stack: st,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	slen := len(this.stack) - 1
	cur := this.stack[slen]
	this.stack = this.stack[:slen]
	val := cur.Val
	cur = cur.Right
	for cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/*
226. 翻转二叉树

翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

/*
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//返回的是从根到叶子的路径： 那么需要判断的是叶子街子安为根的时候就将其加入到结果数组中。当左右节点不为空时，递归调用，并且传递path和当前的值
func binaryTreePaths(root *TreeNode) []string {
	var helper func(root *TreeNode, path string)
	res := []string{}
	helper = func(root *TreeNode, path string) {
		if root.Left == nil && root.Right == nil {
			//代表叶子节点：
			res = append(res, fmt.Sprintf("%s%d", path, root.Val))
		}
		if root.Left != nil {
			helper(root.Left, fmt.Sprintf("%s%d->", path, root.Val))
		}
		if root.Right != nil {
			helper(root.Right, fmt.Sprintf("%s%d->", path, root.Val))
		}
	}
	if root != nil {
		helper(root, "")
	}
	return res
}

/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//二叉搜索树按中序遍历是一个有序数组
func kthSmallest(root *TreeNode, k int) int {

	return 0
}

/*
1026. 节点与其祖先之间的最大差值
给定二叉树的根节点 root，找出存在于不同节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。

（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）

 

输入：[8,3,10,1,6,null,14,null,null,4,7,13]
输出：7
解释：
我们有大量的节点与其祖先的差值，其中一些如下：
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * 解题： 只需要找到左右俩边的最大值和最小值即可
 */
func maxAncestorDiff(root *TreeNode) int {
	var (
		helper func(root *TreeNode, maxv int, minv int) int
	)

	if root == nil {
		return 0
	}

	helper = func(root *TreeNode, maxv int, minv int) int {
		if root == nil {
			return maxv - minv
		}
		if root != nil {
			if root.Val < minv {
				minv = root.Val
			}
			if root.Val > maxv {
				maxv = root.Val
			}
		}

		return int(math.Max(float64(helper(root.Left, maxv, minv)), float64(helper(root.Right, maxv, minv))))
	}
	return helper(root, root.Val, root.Val)

}

/*
1022. 从根到叶的二进制数之和


给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。

对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。

以 10^9 + 7 为模，返回这些数字之和。

 

示例：



输入：[1,0,1,0,1,0,1]
输出：22
解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
//需要计算每次移动时的val的值，然后如果是叶子节点，则计算返回值，然后将左右子树的值相加。
func sumRootToLeaf(root *TreeNode) int {
	var (
		helper func(root *TreeNode, sum int)
		ret    int
	)

	helper = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum*2 + root.Val
		//或者使用sum = sum <<1 ; sum +=root.Val
		if root.Left == nil && root.Right == nil {
			ret = ret + sum

		}
		helper(root.Left, sum)
		helper(root.Right, sum)
	}
	helper(root, 0)
	return ret
}

/*
1008. 先序遍历构造二叉树

返回与给定先序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。

(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，而 node.right 的任何后代，值总 > node.val。此外，先序遍历首先显示节点的值，然后遍历 node.left，接着遍历 node.right。）

 

示例：

输入：[8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */
func bstFromPreorder(preorder []int) *TreeNode {
	//先序遍历的二叉搜索树，首先分清楚，左右俩个节点的分界线为preorder[0]
	//然后分别构造左右子树
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	piviot := 0
	//查找左右子树的分界点
	for i := 0; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			piviot = i
			break
		}
	}

	left := preorder[:piviot]
	right := preorder[piviot:]

	for i,j:=0,piviot;i +j <len(preorder); {
		if preorder[i] <
	}
}

func main() {
	//ret := generateTrees(2)
	//fmt.Println(len(ret))

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(maxAncestorDiff(root))

	fmt.Println(sumRootToLeaf(root))
	//fmt.Println(binaryTreePaths(root))

	//fmt.Println(levelOrderBottom(root))
	//fmt.Println(zigzagLevelOrder(root))

	//inorder := []int{9, 3, 15, 20, 7}
	//postorder := []int{9, 15, 7, 20, 3}
	//consuctTree(inorder, postorder)

	//fmt.Println(pathSum(root, 7))
	//
	//root1 := &Node{val: 1}
	//root1.left = &Node{val: 2}
	//root1.right = &Node{val: 3}
	//root1.left.left = &Node{val: 4}
	//root1.left.right = &Node{val: 5}
	//root1.right.left = &Node{val: 6}
	//root1.right.left = &Node{val: 7}
	//connect(root1)

}
