package structure

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"math"
)

/*
参考： https://blog.csdn.net/luckyxiaoqiang/article/details/7518888
      https://annatarhe.github.io/2016/12/22/btree-search.html
      https://segmentfault.com/a/1190000008850005
	  https://segmentfault.com/a/1190000008760308

*/
type Node struct {
	Value        int
	lNode, rNode *Node
}

func (n *Node) Print() {
	if n != nil {
		fmt.Printf("%v\n", n.Value)
	}
}

func (n *Node) SetValue(v int) {
	if n == nil {
		return
	}
	n.Value = v
}

/**
前序遍历：先遍历根节点，然后遍历左节点，然后遍历右节点
 */
func (n *Node) PreOrder() {
	if (n == nil) {
		return
	}
	n.Print()
	n.lNode.PreOrder()
	n.rNode.PreOrder()
}

/*
中序遍历：
	先遍历左节点，然后访问根节点，然后访问右节点
 */
func (n *Node) MidOrder() {
	if n == nil {
		return
	}
	n.lNode.MidOrder()
	n.Print()
	n.rNode.MidOrder()
}

/**
后序遍历：
左--->右---> 根
 */
func (n *Node) PostOrder() {
	if n == nil {
		return
	}
	n.lNode.PostOrder()
	n.rNode.PostOrder()
	n.Print()
}

func CreateNode(v int) *Node {
	return &Node{Value: v}
}

/**
层次遍历：
	从第一层开始，然后第二层，第三层遍历

相当于广度优先搜索，
	使用队列实现。队列初始化，将根节点压入队列。
 	当队列不为空，进行如下操作：
		弹出一个节点，访问，若左子节点或右子节点不为空，将其压入队列。
 */
func (n *Node) LevelOrder() {
	if n == nil {
		return
	}
	q := queue.Queue{}
	q.Enqueue(n)

	for q.Len() != 0 {
		p := (q.Dequeue()).(*Node)
		if p.lNode != nil {
			q.Enqueue(p.lNode)
		}

		if p.rNode != nil {
			q.Enqueue(p.rNode)
		}
		p.Print()
	}
}

/**
将二叉树变为有序的双向链表
要求不能创建新节点，只调整指针。
递归解法：
（1）如果二叉树查找树为空，不需要转换，对应双向链表的第一个节点是NULL，最后一个节点是NULL
（2）如果二叉查找树不为空：
	如果左子树为空，对应双向有序链表的第一个节点是根节点，左边不需要其他操作；
	如果左子树不为空，转换左子树，二叉查找树对应双向有序链表的第一个节点就是左子树转换后双向有序链表的第一个节点，同时将根节点和左子树转换后的双向有序链 表的最后一个节点连接；
	如果右子树为空，对应双向有序链表的最后一个节点是根节点，右边不需要其他操作；
	如果右子树不为空，对应双向有序链表的最后一个节点就是右子树转换后双向有序链表的最后一个节点，同时将根节点和右子树转换后的双向有序链表的第一个节点连 接。
 */

func (n *Node) Convert() {

}

/*
1. 求二叉树中的节点个数
递归解法：
（1）如果二叉树为空，节点个数为0
（2）如果二叉树不为空，二叉树节点个数 = 左子树节点个数 + 右子树节点个数 + 1
 */
func (n *Node) GetNodeNum() int {
	if n == nil {
		return 0
	}
	lnums := n.lNode.GetNodeNum()
	rnums := n.rNode.GetNodeNum()
	return lnums + rnums + 1
}

/*
2. 求二叉树中叶子节点的个数
递归解法：
（1）如果二叉树为空，返回0
（2）如果二叉树不为空且左右子树为空，返回1
（3）如果二叉树不为空，且左右子树不同时为空，返回左子树中叶子节点个数加上右子树中叶子节点个数
 */

func (n *Node) GetLeafNodeNum() int {
	if n == nil {
		return 0
	}
	if n.lNode == nil && n.rNode == nil {
		return 1
	}
	return n.lNode.GetNodeNum() + n.rNode.GetLeafNodeNum()
}

/*
3. 求树的深度:
如果二叉树为空，返回0
不为空： 二叉树的深度=max(左树深度，右树深度)+1
 */
func (n *Node) GetDepth() int {
	if n == nil {
		return 0
	}
	lnum := n.lNode.GetDepth()
	rnum := n.rNode.GetDepth()
	return int(math.Max(float64(lnum), float64(rnum))) + 1
}

/**
4. 二叉树第K层的节点个数
递归解法：
（1）如果二叉树为空或者k<1返回0
（2）如果二叉树不为空并且k==1，返回1
（3）如果二叉树不为空且k>1，返回左子树中k-1层的节点个数与右子树k-1层节点个数之和
 */

func (n *Node) GetNodeNumKLevel(k int) int {
	if n == nil || k < 1 {
		return 0
	}
	if k == 1 {
		return 1
	}

	return n.lNode.GetNodeNumKLevel(k-1) + n.rNode.GetNodeNumKLevel(k-1)
}

/*
func main() {
	root := Node{Value: 3}
	root.lNode = &Node{}
	root.rNode = &Node{5, nil, nil}
	root.rNode.lNode = new(Node)
	root.rNode.lNode.SetValue(4)
	root.lNode.rNode = CreateNode(2)
	root.lNode.lNode = CreateNode(6)
	fmt.Println("前序遍历: ")
	root.PreOrder()
	fmt.Println()
	fmt.Println("中序遍历: ")
	root.MidOrder()
	fmt.Println()
	fmt.Println("后序遍历: ")
	root.PostOrder()
	fmt.Println("层次遍历: ")
	root.LevelOrder()

	fmt.Printf("node数量：  %v \n",root.GetNodeNum())
	fmt.Printf("叶子节点个数：  %d \n", root.GetLeafNodeNum())
	fmt.Printf("二叉树深度:  %d\n",root.GetDepth())
	fmt.Printf("二叉树第K层节点个数：  %d", root.GetNodeNumKLevel(3))
}
*/

/*
5. 判断两棵二叉树是否结构相同
不考虑数据内容。结构相同意味着对应的左子树和对应的右子树都结构相同。
递归解法：
（1）如果两棵二叉树都为空，返回真
（2）如果两棵二叉树一棵为空，另一棵不为空，返回假
（3）如果两棵二叉树都不为空，如果对应的左子树和右子树都同构返回真，其他返回假

 */

func (n *Node) StructureCmp(n1 *Node) bool {
	if n == nil && n1 == nil {
		return true
	} else if n == nil || n1 == nil {
		return false
	}
	bleft := n.lNode.StructureCmp(n1.lNode)
	bright := n.rNode.StructureCmp(n1.rNode)
	return bleft && bright
}

/*
6. 求二叉树的镜像: 对于每个节点，我们交换它的左右孩子即可。
递归解法：
（1）如果二叉树为空，返回空
（2）如果二叉树不为空，求左子树和右子树的镜像，然后交换左子树和右子树

 */
func (n *Node) Mirror() {
	if n == nil {
		return
	}
	temp := n.lNode
	n.lNode = n.rNode
	n.rNode = temp
	n.lNode.Mirror()
	n.rNode.Mirror()
}



/*
7. 判断二叉树是不是平衡二叉树
递归解法：
（1）如果二叉树为空，返回真
（2）如果二叉树不为空，如果左子树和右子树都是AVL树并且左子树和右子树高度相差不大于1，返回真，其他返回假
平衡二叉树： 左右子树深度相差不为1

 */
func (n *Node) IsAVL(height int) bool {
	if n == nil {
		height = 0
		return true
	}
	var lheigth, rheight int
	resleft := n.lNode.IsAVL(lheigth)
	resright := n.rNode.IsAVL(rheight)

	temp := float64(lheigth - rheight)
	if resleft && resright && int(math.Abs(temp)) <= 1 {
		height = int(math.Max(float64(lheigth), float64(rheight)) + 1)
		return true
	} else {
		height = int(math.Max(float64(lheigth), float64(rheight)) + 1)
		return false
	}
}

/*
8. 求二叉树中两个节点的最低公共祖先节点
递归解法：
（1）如果两个节点分别在根节点的左子树和右子树，则返回根节点
（2）如果两个节点都在左子树，则递归处理左子树；如果两个节点都在右子树，则递归处理右子树
参考代码如下：
 */

func (n *Node) FindNode(n1 *Node) {

}

/*
9. 求二叉树中节点的最大距离
即二叉树中相距最远的两个节点之间的距离。
递归解法：
（1）如果二叉树为空，返回0，同时记录左子树和右子树的深度，都为0
（2）如果二叉树不为空，最大距离要么是左子树中的最大距离，要么是右子树中的最大距离，要么是左子树节点中到根节点的最大距离+右子树节点中到根节点的最大距离，同时记录左子树和右子树节点中到根节点的最大距离。
 */

func (n *Node) GetMaxDistance(maxleft int, maxright int) int {
	var maxLeft, maxRight int
	if n == nil {
		maxLeft = 0;
		maxRight = 0;
		return 0
	}

	var maxDistLeft, maxDistRight int
	var maxLL, maxLR, maxRL, maxRR int

	if n.lNode != nil {
		maxDistLeft = n.lNode.GetMaxDistance(maxLL, maxLR)
		maxLeft = int(math.Max(float64(maxLL), float64(maxLR)) + 1)
	} else {
		maxDistLeft = 0
		maxLeft = 0
	}

	if n.rNode != nil {
		maxDistRight = n.rNode.GetMaxDistance(maxRL, maxRR)
		maxRight = int(math.Max(float64(maxRL), float64(maxRR)) + 1)
	} else {
		maxDistRight = 0
		maxRight = 0
	}
	y := float64(math.Max(float64(maxDistRight), float64(maxDistLeft)))
	f := float64(maxLeft + maxRight)
	return int(math.Max(y, f))

}


func main() {
	root := Node{Value: 3}
	root.lNode = &Node{}
	root.rNode = &Node{5, nil, nil}
	root.rNode.lNode = new(Node)
	root.rNode.lNode.SetValue(4)
	root.lNode.rNode = CreateNode(2)
	root.lNode.lNode = CreateNode(6)

	right := &Node{Value: 3}
	right.lNode = &Node{}
	right.rNode = &Node{5, nil, nil}
	right.rNode.lNode = new(Node)
	right.rNode.lNode.SetValue(4)
	right.lNode.rNode = CreateNode(2)
	right.lNode.lNode = CreateNode(6)

	fmt.Println("俩个树结构是否相同：   %v", root.StructureCmp(right))

	fmt.Println("镜像前======")
	root.LevelOrder()
	fmt.Println("镜像后=====")
	root.Mirror()
	root.LevelOrder()
	fmt.Println("=========")
	fmt.Printf("节点之间的最大距离：     %d\n",right.GetMaxDistance(0,3))
}
/*
10. 由前序遍历序列和中序遍历序列重建二叉树
二叉树前序遍历序列中，第一个元素总是树的根节点的值。中序遍历序列中，左子树的节点的值位于根节点的值的左边，右子树的节点的值位
于根节点的值的右边。
递归解法：
（1）如果前序遍历为空或中序遍历为空或节点个数小于等于0，返回NULL。
（2）创建根节点。前序遍历的第一个数据就是根节点的数据，在中序遍历中找到根节点的位置，可分别得知左子树和右子树的前序和中序遍
历序列，重建左右子树。
同样，有中序遍历序列和后序遍历序列，类似的方法可重建二叉树，但前序遍历序列和后序遍历序列不同恢复一棵二叉树，证明略。

 */
 func (n *Node)RebuildBinaryTree(){

 }
/*
判断二叉树是不是完全二叉树
若设二叉树的深度为h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全
二叉树。
有如下算法，按层次（从上到下，从左到右）遍历二叉树，当遇到一个节点的左子树为空时，则该节点右子树必须为空，且后面遍历的节点左
右子树都必须为空，否则不是完全二叉树。

 */

/*
func main() {
	root := Node{Value: 3}
	root.lNode = &Node{}
	root.rNode = &Node{5, nil, nil}
	root.rNode.lNode = new(Node)
	root.rNode.lNode.SetValue(4)
	root.lNode.rNode = CreateNode(2)
	fmt.Println("前序遍历: ")
	root.PreOrder()
	fmt.Println()
	fmt.Println("中序遍历: ")
	root.MidOrder()
	fmt.Println()
	fmt.Println("后序遍历: ")
	root.PostOrder()

}
*/
