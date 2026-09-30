/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	subRootVals := []string{}
	rootVals := []string{}
	dfs(subRoot, &subRootVals)
	dfs(root, &rootVals)

	subStr := "," + strings.Join(subRootVals, ",") + ","
	rootStr := "," + strings.Join(rootVals, ",") + ","

	return strings.Contains(rootStr, subStr)
}

func dfs(node *TreeNode, vals *[]string) {
	if node == nil {
		*vals = append(*vals, "#")
		return
	}

	*vals = append(*vals, fmt.Sprintf("%d", node.Val))
	dfs(node.Left, vals)
	dfs(node.Right, vals)
}