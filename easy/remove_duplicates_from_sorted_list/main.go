package main 

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	l := head
	for l.Next != nil {
		if l.Val == l.Next.Val {	
			l.Next = l.Next.Next
		} else {
			l = l.Next		
		}	
	}
	return head
}

func main() {
	// head := &ListNode{
	// 	Val: 1, Next: &ListNode{
	// 		Val: 1, Next: &ListNode{
	// 			Val:  2, Next: nil,
	// 		},
	// 	},
	// }

	head := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val:  2, Next: &ListNode{
					Val:  3, Next: &ListNode{
						Val:  3, Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(deleteDuplicates(head))
}