package main


import (
	"sync"
	"fmt"
)

//二叉搜索树
type BinaryNode struct {
	key   int         // 中序遍历的节点序号
	value int         //节点存储的值
	left  *BinaryNode //左子节点
	right *BinaryNode //右子节点
}

// ItemBinarySearchTree the binary search tree of Items
type ItemBinarySearchTree struct {
	root *BinaryNode
	lock sync.RWMutex
}

/*
插入操作需要使用到递归，插入操作需要从上到下查找新节点在树中合适的位置：新节点的值小于任意节点，则向左子树继续寻找，同理向右子树查找，直到查找到树叶节点再插入
遍历操作有三种方式：
// Insert inserts the Item t in the tree
 */
func (b *ItemBinarySearchTree) Insert(key int, v int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	newNode := &BinaryNode{key, v, nil, nil}
	if b.root == nil {
		b.root = newNode
	} else {
		insertNode(b.root, newNode)
	}
}

/**
如果node的key 小于新节点的key，在左节点中查找然后排序
如果node的key小于新节点的key，在右节点中查找然后插入
 */
func insertNode(node, newNode *BinaryNode) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

/**
因为已经排好序，所以只需要比较节点的值即可，如果小于则左节点排序，如果大于则右节点排序
 */
func (b *ItemBinarySearchTree) Search(key int) bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	if b.root == nil {
		return false
	}
	return search(b.root, key)
}

func search(n *BinaryNode, key int) bool {
	if n == nil {
		return false
	}
	if n.key < key {
		search(n.right, key)
	} else if (n.key > key) {
		search(n.left, key)
	}
	return true
}

/*
也是先按照左右节点查找然后再删除
 */
func (b *ItemBinarySearchTree) Delete(key int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.root == nil {
		return
	}
	delete(b.root, key)
}

func delete(b *BinaryNode, key int) *BinaryNode {
	if b == nil {
		return nil
	}
	if b.key < key {
		b.right = delete(b.right, key)
		return b.right
	} else if b.key > key {
		b.left = delete(b.left, key)
		return b.left
	}

	// if key == node.key 则需要处理node left和right的节点上移问题
	if b.left == nil && b.right == nil {
		b = nil
		return b
	}
	if b.left == nil {
		b = b.right
		return b
	}
	if b.right == nil {
		b = b.left
		return b
	}
	//如果有多个节点的话，删除后的树仍然需要满足二叉搜索树的原则： 左小右大的原则
	//用右子树的最左节点替代当前节点
	//用左子树最右的节点替代当前节点

	//// 要删除的节点有 2 个子节点，找到右子树的最左节点，替换当前节点
	mostLeftNode := b.right

	for {
		if mostLeftNode != nil && mostLeftNode.left != nil {
			mostLeftNode = mostLeftNode.left
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	b.key, b.value = mostLeftNode.key, mostLeftNode.value
	b.right = delete(b.right, b.key)
	return b

}

/**
一直遍历到最左的节点
 */
func (b *ItemBinarySearchTree) Min() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	if b.root == nil {
		return -1
	}
	node := b.root
	for {
		if node.left == nil {
			return node.value
		}
		node = node.left
	}
}

/**
一直找到最右的节点
 */
func (b *ItemBinarySearchTree) Max() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	if b.root == nil {
		return -1
	}
	node := b.root
	for {
		if node.right == nil {
			return node.value
		}
		node = node.right
	}
}

/*
中序遍历（in-order）：左子树–>根结点–> 右树：1->2->3->4->5->6->7->8->9->10->11
 */
func (b *ItemBinarySearchTree) InOrder(printFunc func(v int)) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	inorder(b.root, printFunc)
}

func inorder(b *BinaryNode, printFunc func(v int)) {
	if b == nil {
		return
	}
	inorder(b.left, printFunc)
	printFunc(b.value)
	inorder(b.right, printFunc)
}

/*
先序遍历（pre-order）：根结点–>左子树–>右子树：8->4->2->1->3->6->5->7 >10->9->11
 */
func (b *ItemBinarySearchTree) PreOrder(printFunc func(value int)) {
	if b.root == nil {
		return
	}
	preOrder(b.root, printFunc)
}

func preOrder(b *BinaryNode, printFunc func(value int)) {
	if b == nil {
		return
	}
	printFunc(b.value)
	preOrder(b.left, printFunc)
	preOrder(b.right, printFunc)
}

/*
后序遍历（post-order）：左子树–>右子树–>根结点：1->3->2->5->7->6->4->9->11->10->8
 */
func (b *ItemBinarySearchTree) PostOrder(printFunc func(value int)) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	postOrder(b.root,printFunc)
}

func postOrder(b *BinaryNode, printFunc func(value int)) {
	if b == nil {
		return
	}
	postOrder(b.left, printFunc)
	postOrder(b.right, printFunc)
	printFunc(b.value)
}


// String prints a visual representation of the tree
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *BinaryNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}