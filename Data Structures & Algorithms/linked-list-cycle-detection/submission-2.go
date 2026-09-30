/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    var nullPointer *ListNode
    if head == nullPointer{
        return false
    }
    slow := head
    fast := head.Next
    
    if fast == nullPointer{
        return false
    }
    

    for slow != fast {
        slow = slow.Next
        fast = fast.Next
        if fast == nullPointer{
            return false
        }
        fast = fast.Next
        if fast == nullPointer{
            return false
        }
    }

    return true

}
