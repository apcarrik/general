# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev: LLNode = None
        current: LLNode = head
        while current is not None:
            next: LLNode = current.next
            current.next = prev
            prev = current
            current = next
        return prev
