# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if not list1 and not list2:
            return None
        elif not list1 or not list2:
            return list1 or list2
        
        head = list1 if list1.val <= list2.val else list2

        running = head
        lane = running.next
        alternate_lane = list2 if list1 is head else list1

        while lane is not None and alternate_lane is not None:
            while lane is not None and lane.val < alternate_lane.val:
                running = lane
                lane = running.next
            
            temp = lane
            lane = alternate_lane.next
            running.next = alternate_lane
            running = running.next
            alternate_lane = temp

        if lane is None:
            running.next = alternate_lane

        return head

