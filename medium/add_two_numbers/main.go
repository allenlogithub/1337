package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	current := head
	carry := 0 
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}		
		
		carry = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next		
	}
	
	fmt.Println("=====")
	fmt.Println(head, head.Next, head.Next.Next, head.Next.Next.Next)
	return head.Next
}

func main() {
	l1 := &ListNode{
		Val: 2, Next: &ListNode{
			Val: 4, Next: &ListNode{
				Val:  3, Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5, Next: &ListNode{
			Val: 6, Next: &ListNode{
				Val:  4, Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}