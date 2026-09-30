# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    result = 0
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return self.result
        self.maxDepthHelper(root)
        return self.result

    def maxDepthHelper(self, node, depth=1):
        if not node:
            return
        self.maxDepthHelper(node.left, depth=depth+1)
        self.result = max(self.result, depth)
        self.maxDepthHelper(node.right, depth=depth+1)