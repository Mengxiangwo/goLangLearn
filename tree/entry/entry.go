package main

import (
	"fmt"
	"goLangLearn/tree"
	"golang.org/x/tools/container/intsets"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(10000)
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(100000))
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateNode(2)
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)

	fmt.Println(root)
	root.Traverse()

	myNode := myTreeNode{&root}
	myNode.postOrder()

	testSparse()
	/*root.Right.Left.Print()

	root.SetValue(100)
	root.Print()

	var pRoot *Node
	fmt.Println(pRoot)
	pRoot.SetValue(200)
	//pRoot.Print() //nil调用Value
	fmt.Println(pRoot)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	root.Print()

	fmt.Println(root)

	nodes := []Node{
		{Value:3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)*/
}
