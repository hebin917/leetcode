package main

import (
	"math"
	"github.com/eapache/queue"
	"fmt"
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
	fmt.Println("heigt:  ",h)
	res := [][]int{}
	for i := h; i >=1 ; i-- {
		tmp := []int{}
		getGivenLevel(root, i,&tmp)
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

func main() {
	//ret := generateTrees(2)
	//fmt.Println(len(ret))

	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(levelOrderBottom(root))

}
