package main

import "fmt"

func main() {
	fmt.Println("树的一些结构")
	//+++++++++++++++++++++++++++++++++++++++++++++++//
	// 普通二叉树练习
	//                 10
	//                /  \
	//               /    \
	//              6      4
	//             / \    /
	//            1   5  2

	// 创建二叉树 思路：先创建节点，然后把节点连接起来
	root := newTree(10) // 创建根节点10
	root.addLeft(6).addRight(4)
	{
		// 以根节点的左节点为根节点创建一个二叉树  根6 左1 右5
		root.left.addLeft(1).addRight(5)
		// 以根节点的右节点为根节点创建一个二叉树 根4 左2 右空
		root.right.addLeft(2)
	}
	root.PreOrder()

}

type tree struct {
	node  int
	left  *tree
	right *tree
}

func newTree(node int) *tree {
	return &tree{node: node}
}

func (t *tree) addLeft(treenode interface{}) *tree {
	if tree, ok := treenode.(*tree); ok {
		t.left = tree
	}
	if node, ok := treenode.(int); ok {
		t.left = newTree(node)
	}
	return t
}

func (t *tree) addRight(treenode interface{}) *tree {
	if tree, ok := treenode.(*tree); ok {
		t.right = tree
	}
	if node, ok := treenode.(int); ok {
		t.right = newTree(node)
	}
	return t
}

func (this *tree) PreOrder() {
	if this == nil {
		return
	}
	fmt.Printf("---->%d", this.node)
	this.left.PreOrder()
	this.right.PreOrder()
}
