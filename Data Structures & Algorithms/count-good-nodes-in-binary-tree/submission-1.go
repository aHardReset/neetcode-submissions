/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	happyNodes := 0
	maxPathVal := root.Val - 1
	goodNodesHelper(root, &happyNodes, maxPathVal)
	return happyNodes
    
}

func goodNodesHelper(node *TreeNode, happyNodes *int, maxPathVal int) {
	if node == nil {
		return
	}
	if node.Val >= maxPathVal {
		*happyNodes = *happyNodes + 1
	}

	goodNodesHelper(node.Left, happyNodes, max(node.Val, maxPathVal))
	goodNodesHelper(node.Right, happyNodes, max(node.Val, maxPathVal))

}