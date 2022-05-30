package main

import (
	"fmt"
)

//int类型二叉树
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

func (t *TreeNode) setValue(value int) {
	t.Value = value
}

func (t TreeNode) print() {
	fmt.Println(t.Value)
}

func (t *TreeNode) traverse() {
	if t.Left != nil {
		t.Left.traverse()
	}
	t.print()
	if t.Right != nil {
		t.Right.traverse()
	}
}

func (t *TreeNode) preTraverse() {
	if t == nil {
		return
	}
	t.print()
	if t.Left != nil {
		t.Left.preTraverse()
	}

	if t.Right != nil {
		t.Right.preTraverse()
	}
}

func (t *TreeNode) TraverseFunc(f func(*TreeNode))  {
	if t == nil {
		return
	}

	t.Left.TraverseFunc(f)
	f(t)
	t.Right.TraverseFunc(f)
}

func (t *TreeNode) traverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		t.TraverseFunc(func(node *TreeNode) {
			out <- node
		})
		close(out)
	}()

	return out
}

func main() {
	var root *TreeNode

	root = &TreeNode{
		Value: 3,
	}

	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Left.Right = CreateTreeNode(2)
	root.Left.Right.Left = &TreeNode{1, nil, nil}
	root.Left.Right.Left.Left = &TreeNode{8, nil, nil}
	root.Left.Right.Left.Right = &TreeNode{9, nil, nil}
	root.Left.Right.Right = &TreeNode{6, nil, nil}
	root.Left.Right.Right.Left = &TreeNode{7, nil, nil}
	root.Right.Left.setValue(4)
	//root.preTraverse()
	//
	nodeCount := 0
	root.TraverseFunc(func(node *TreeNode) {
		nodeCount++
	})

	//把treeNode中最大的数取出
	c := root.traverseWithChannel()
	maxNode := 0
	for node := range c{
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max Node Value:", maxNode)


	//data := []int{1,2,4,8,13,30,5}
	//
	//root := CreateTreeNode(1)
	//
	//for _, v := range data{
	//	root.insertIntoBST(v)
	//}
	//
	//fmt.Println(root)
}

func isSymmetric(root *TreeNode)  {

}

// 判断二叉搜索树是否对称
func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil{
		return false
	}
	return left.Value == right.Value && compare(left.Right, right.Left) && compare(left.Left, right.Right)
}

