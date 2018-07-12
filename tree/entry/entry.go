package main

import (
	"fmt"
	"goLangLearn/tree"
)

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
