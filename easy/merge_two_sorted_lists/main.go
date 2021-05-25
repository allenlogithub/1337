package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	node = head
    fmt.Println(node)

	for {
        // fmt.Println(l1, l2)
		if l1 == nil {
			node.Next = l2
			break
		}
		if l2 == nil {
			node.Next = l1
            break
		}
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
		fmt.Println(node)
	}
	fmt.Println(node)
	return head
}

func main() {
	l1 := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists(l1, l2)
}
