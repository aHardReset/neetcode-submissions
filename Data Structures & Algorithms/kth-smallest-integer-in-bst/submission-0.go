/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	smallCounter := [3]int{0, 0, k}  // value, currentK, k
	kthSmallestHelper(root, &smallCounter)
	return smallCounter[0]
}

func kthSmallestHelper(node *TreeNode, smallCounter *[3]int) {
	if node == nil || smallCounter[1] == smallCounter[2] {
		return
	}
	kthSmallestHelper(node.Left, smallCounter)
	smallCounter[1] = smallCounter[1] + 1
	if smallCounter[1] == smallCounter[2] {
		smallCounter[0] = node.Val
	}
	kthSmallestHelper(node.Right, smallCounter)
}