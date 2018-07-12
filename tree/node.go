package tree

import "fmt"

//结构体
type Node struct {
	Value int
	Left, Right *Node
}

//Node的接收者，Node的方法
func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node value is nil")

		return
	}
	node.Value = value
}

//工厂函数,Node的__construct
func CreateNode(value int) *Node {
	return &Node{Value:value}
}