package alg

import (
	"fmt"
)

func PreOrderRecursion(n *TreeNode) {
	if n != nil {
		fmt.Println(n.Val)
		PreOrderRecursion(n.Left)
		PreOrderRecursion(n.Right)
	}
}

func InOrderRecursion(n *TreeNode) {
	if n != nil {
		InOrderRecursion(n.Left)
		fmt.Println(n.Val)
		InOrderRecursion(n.Right)
	}
}

func LaterOrderRecursion(n *TreeNode) {
	if n != nil {
		LaterOrderRecursion(n.Left)
		LaterOrderRecursion(n.Right)
		fmt.Println(n.Val)
	}
}

func PreOrder(n *TreeNode) {
	if n != nil {
		stack := []*TreeNode{}
		stack = append(stack, n)
		for len(stack) != 0 {
			head := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(head.Val)
			if head.Right != nil {
				stack = append(stack, head.Right)
			}
			if head.Left != nil {
				stack = append(stack, head.Left)
			}
		}
	}
}

func InOrder(n *TreeNode) {
	if n != nil {
		stack := []*TreeNode{}
		for len(stack) != 0 || n != nil {
			if n != nil {
				stack = append(stack, n)
				n = n.Left
			} else {
				n = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Println(n.Val)
				n = n.Right
			}
		}
	}
}

func LaterOrder(n *TreeNode) {
	if n != nil {
		stack1 := []*TreeNode{}
		stack2 := []*TreeNode{}
		stack1 = append(stack1, n)
		for len(stack1) != 0 {
			head := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			stack2 = append(stack2, head)
			if head.Left != nil {
				stack1 = append(stack1, head.Left)
			}
			if head.Right != nil {
				stack1 = append(stack1, head.Right)
			}
		}
		for len(stack2) != 0 {
			fmt.Println(stack2[len(stack2)-1].Val)
			stack2 = stack2[:len(stack2)-1]
		}
	}
}
