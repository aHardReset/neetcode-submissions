/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	subRootChan := make(chan []string)
    rootChan := make(chan []string)
    
    go func() {
        vals := []string{}
        dfs(subRoot, &vals)
        subRootChan <- vals
    }()
    
    go func() {
        vals := []string{}
        dfs(root, &vals)
        rootChan <- vals
    }()

	subRootVals := <-subRootChan
    rootVals := <-rootChan

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