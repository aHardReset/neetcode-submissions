/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    isBalanced := true
    isBalancedHelper(root, &isBalanced)
    return isBalanced
    
}

func isBalancedHelper(node *TreeNode, isBalanced *bool) int {
    if node == nil || *isBalanced == false {
        return 0
    }

    left := isBalancedHelper(node.Left, isBalanced)
    right := isBalancedHelper(node.Right, isBalanced)
    if (left - right) > 1 || (left - right) < -1{
        *isBalanced = false
    }
    return max(left, right) + 1



}
