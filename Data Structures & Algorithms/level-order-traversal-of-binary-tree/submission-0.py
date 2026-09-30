# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from collections import deque 

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        q = deque()
        level = 0
        if root:
            q.append((level, root))
        level_nodes = []
        level_order = []
        while q:
            current_level, current_node = q.popleft()

            if current_level != level:
                level_order.append(level_nodes)
                level_nodes = []
                level = current_level

            level_nodes.append(current_node.val)
            
            if current_node.left is not None:
                q.append((level + 1, current_node.left))
            if current_node.right is not None:
                q.append((level + 1, current_node.right))
        if level_nodes:
            level_order.append(level_nodes)
        return level_order

