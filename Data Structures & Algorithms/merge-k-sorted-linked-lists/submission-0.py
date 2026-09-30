# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:    
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        if not lists:
            return None

        linked_lists = []
        for l in lists:
            current = l
            while current:
                linked_lists.append(current)
                current = current.next

        def get_val(node):
            return node.val

        linked_lists = sorted(linked_lists, key=get_val)
        for i in range(len(linked_lists)-1):
            linked_lists[i].next = linked_lists[i+1]
        linked_lists[len(linked_lists)-1].next = None
        return linked_lists[0]
        