package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	value int
	lchild *TreeNode
	rchild *TreeNode
}

// 创建节点
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{value: v}
}

// 打印
func (this *TreeNode) Print() {
	fmt.Print(this.value, " ")
}

// 前序遍历
func (this *TreeNode) PreOrderTraverse() {
	if nil == this {
		return
	}
	this.Print()
	this.lchild.PreOrderTraverse()
	this.rchild.PreOrderTraverse()
}

// 中序遍历
func (this *TreeNode) InOrderTraverse() {
	if nil == this {
		return
	}
	this.lchild.InOrderTraverse()
	this.Print()
	this.rchild.InOrderTraverse()
}

// 后序遍历
func (this *TreeNode) PostOrderTraverse() {
	if nil == this {
		return
	}
	this.lchild.PostOrderTraverse()
	this.rchild.PostOrderTraverse()
	this.Print()
}

// 设置节点值
func (this *TreeNode) SetValue(value int) {
	if nil == this {
		fmt.Println("setting value to nil")
		return
	}
	this.value = value
}

// 求二叉树深度、高度、层
func (this *TreeNode) Depth() int {
	if nil == this {
		return 0
	}
	lDepth := 1 + this.lchild.Depth()
	rDepth := 1 + this.rchild.Depth()
	return int(math.Max(float64(lDepth), float64(rDepth)))
}

func main() {
	root := TreeNode{value: 3}
	root.lchild = &TreeNode{}
	root.rchild = &TreeNode{5, nil, nil}
	root.rchild.lchild = new(TreeNode)
	root.rchild.lchild.SetValue(4)
	root.lchild.rchild = NewTreeNode(2)
	fmt.Print("前序遍历: ")
  root.PreOrderTraverse()
  fmt.Println()
  fmt.Print("中序遍历: ")
  root.InOrderTraverse()
  fmt.Println()
  fmt.Print("后序遍历: ")
	root.PostOrderTraverse()
	fmt.Println()
	fmt.Print("最大深度: ")
	treeDepth := root.Depth()
	fmt.Println(treeDepth)
}
