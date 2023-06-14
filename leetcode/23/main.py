# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap: list[tuple[int,int]] = []
        heapq.heapify(heap)
        ret_head: ListNode = None
        ret_tail: ListNode = ret_head

        # initialize heap with smallest element from each list
        for i in range(len(lists)):
            if lists[i] is not None:
                heapq.heappush(heap, (lists[i].val, i))
                lists[i] = lists[i].next
        
        # pop smallest from min heap and replace with next value from same list until all lists are exhausted
        while len(heap) > 0:
            val, i = heapq.heappop(heap)
            # add value to return list
            new_node: ListNode = ListNode(val=val)
            if ret_head is None:
                ret_head = new_node
                ret_tail = new_node
            else:
                ret_tail.next = new_node
                ret_tail = ret_tail.next

            # place new val on heap if applicable
            if lists[i] is not None:
                heapq.heappush(heap, (lists[i].val, i))
                lists[i] = lists[i].next
        return ret_head
