/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func maxDepth(root *TreeNode) int {
    maxDepth := 0
    maxDepthHelper(root, &maxDepth, 0)
    return maxDepth
}

func maxDepthHelper(node *TreeNode, maxDepth *int, depth int) {
    if node == nil {
        return
    }
    currentDepth := depth + 1
    if *maxDepth < currentDepth {
        *maxDepth = currentDepth
    }
    maxDepthHelper(node.Left, maxDepth, currentDepth)
    maxDepthHelper(node.Right, maxDepth, currentDepth)
}