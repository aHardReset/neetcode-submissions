/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    counterP := head
    length := 0
    for counterP != nil {
        length++
        counterP = counterP.Next
    }

    var prev *ListNode
    toRemove := head
    advance := length - n

    for i:=0 ; i<advance; i++ {
        prev = toRemove
        toRemove = toRemove.Next
    }
    next := toRemove.Next

    if prev == nil {
        return head.Next
    } else if next == nil {
        var nilListNode *ListNode
        prev.Next = nilListNode
        return head
    }

    prev.Next = next
    return head
}

