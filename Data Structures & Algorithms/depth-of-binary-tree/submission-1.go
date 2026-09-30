/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func maxDepth(root *TreeNode) int {
    depth := struct{depth int}{depth: 0}
    maxDepthHelper(root, &depth, 0)
    return depth.depth
}

func maxDepthHelper(node *TreeNode, maxDepth *struct{depth int}, depth int) {
    if node == nil {
        return
    }
    currentDepth := depth + 1
    if maxDepth.depth < currentDepth {
        maxDepth.depth = currentDepth
    }
    maxDepthHelper(node.Left, maxDepth, currentDepth)
    maxDepthHelper(node.Right, maxDepth, currentDepth)
}