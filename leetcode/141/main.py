# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# class Solution:
#     def hasCycle(self, head: Optional[ListNode]) -> bool:

class Solution:
  def hasCycle(self, head: ListNode) -> bool:
    fast: ListNode = head
    slow: ListNode = head
    while fast is not None:
      # iterate fast and slow, checking if they point to the same nodes
      if fast.next is None:
        break
      slow = slow.next
      fast = fast.next.next
      if fast == slow:
        return True

    # if fast becomes None, we have reached end of list with no loop
    return False
    
