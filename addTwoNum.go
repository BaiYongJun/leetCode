package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// if carry > 0 {
	// 	tail.Next = &ListNode{Val: carry}
	// }
	return
}

func main() {
	var aa *ListNode
	aa = new(ListNode)
	var l1 = ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 = ListNode{8, &ListNode{9, &ListNode{7, nil}}}
	aa = addTwoNumbers(&l1, &l2)
	fmt.Println(aa.Val)
	fmt.Println(aa.Next.Val)
	fmt.Println(aa.Next.Next.Val)
}
