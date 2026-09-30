/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    diameterOfBinaryTreeHelper(root, &diameter)
    return diameter
    
}

func diameterOfBinaryTreeHelper(node *TreeNode, diameter *int) int {
    if node == nil {
        return 0
    }

    left := diameterOfBinaryTreeHelper(node.Left, diameter)
    right := diameterOfBinaryTreeHelper(node.Right, diameter)
    if (left+right) > *diameter {
        *diameter = left+right
    }
    return max(left,right) + 1

}
