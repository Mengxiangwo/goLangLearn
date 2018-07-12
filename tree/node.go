package main

import "fmt"

//结构体
type treeNode struct {
	value int
	left, right *treeNode
}

//treeNode的接收者，treeNode的方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("node value is nil")

		return
	}
	node.value = value
}

func (node *treeNode) traverse()  {
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

//工厂函数,treeNode的__construct
func createNode(value int) *treeNode {
	return &treeNode{value:value}
}

func main() {
	var root treeNode
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.left.right = createNode(2)
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.right.left.setValue(4)

	fmt.Println(root)
	root.traverse()

	/*root.right.left.print()

	root.setValue(100)
	root.print()

	var pRoot *treeNode
	fmt.Println(pRoot)
	pRoot.setValue(200)
	//pRoot.print() //nil调用value
	fmt.Println(pRoot)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()

	root.print()

	fmt.Println(root)

	nodes := []treeNode{
		{value:3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)*/
}
