/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    indexNodeMap := mapNodes(head)
    headCopy := copyLinkedList(head)
    copyList := listNodes(headCopy)
    assignRandom(head, indexNodeMap, copyList)
    return headCopy
}

func assignRandom(head *Node, indexes map[*Node]int, copyList []*Node) {
    curr := head
    for curr != nil{
        if curr.Random != nil {
            originIdx := indexes[curr]
            targetIdx := indexes[curr.Random]
            copyList[originIdx].Random = copyList[targetIdx]
        }
        curr = curr.Next
    }
}

func mapNodes(head *Node) map[*Node]int {
    nodes := map[*Node]int{}
    curr := head
    i := 0
    for curr != nil {
        nodes[curr] = i
        curr = curr.Next
        i++
    }

    return nodes
}

func listNodes(head *Node) []*Node {
    curr := head
    res := []*Node{}
    for curr != nil {
        res = append(res, curr)
        curr = curr.Next
    }
    return res
}

func copyLinkedList(node *Node) *Node {
    if node == nil{
        return nil
    }
    newNode := &Node{
        Val: node.Val,
        Next: copyLinkedList(node.Next),
    }
    return newNode
}

