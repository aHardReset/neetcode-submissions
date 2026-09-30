# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    good_nodes = 0
    def goodNodes(self, root: TreeNode) -> int:
        self.goodNodesHelper(root, root.val)
        return self.good_nodes
    
    def goodNodesHelper(self, node, max_val_in_path):
        if node is None:
            return
        self.goodNodesHelper(node.left, max(max_val_in_path, node.val))
        if max_val_in_path <= node.val:
            self.good_nodes += 1
        self.goodNodesHelper(node.right, max(max_val_in_path, node.val))