package alg

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//功能：打印节点的值
//参数：nil
//返回值：nil
func (node *TreeNode) Print() {
	fmt.Printf("%d ", node.Val)
}

//功能：设置节点的值
//参数：节点的值
//返回值：nil
func (node *TreeNode) SetValue(value int) {
	node.Val = value
}
func CreateNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func GetTreeNode() *TreeNode {
	root := CreateNode(5)
	root.Left = CreateNode(2)
	root.Right = CreateNode(4)
	root.Left.Right = CreateNode(7)
	root.Left.Right.Left = CreateNode(6)
	root.Right.Left = CreateNode(8)
	root.Right.Right = CreateNode(9)
	return root
}
