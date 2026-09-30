/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    q := []*TreeNode{root} // [nil,nil,nil,nil,nil,nil,ni,nil]
    res := [][]int{} // [[1], [2,3], [4,5,6,7]]
    for len(q) > 0 {
        level := []int{} // []
        nextLevel := []*TreeNode{} // []
        for _, n := range(q){
            if n != nil {
                level = append(level, n.Val)
                nextLevel = append(nextLevel, n.Left)
                nextLevel = append(nextLevel, n.Right)
            }
        }
        if len(level) > 0{
            res = append(res, level)
        }
        q = []*TreeNode{}

        for _, n := range nextLevel{
            q = append(q, n)
        }
    }
    return res
}
