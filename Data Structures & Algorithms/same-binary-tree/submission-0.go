/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    queue := []*TreeNode{}
    queue = append(queue, p)
    queue = append(queue, q)
    for len(queue) > 0 {
        p1 := queue[0]
        q1 := queue[1]
        queue = queue[2:]
        if p1 == nil && q1 == nil{
            continue
        } else if (p1 == nil && q1 != nil) || (p1 != nil && q1 == nil) || (p1.Val != q1.Val) {
            return false
        }

        

        queue = append(queue, p1.Left)
        queue = append(queue, q1.Left)
        queue = append(queue, p1.Right)
        queue = append(queue, q1.Right)

    }
    return true
}
