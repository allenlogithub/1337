/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    fast, slow := head, head
    for fast != nil && fast.Next != nil && slow != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            // return slow
			slow = head
			for {
                if slow == fast {
					return slow
				}
				slow = slow.Next
				fast = fast.Next				
			}
        }
    }       
    return nil
}
