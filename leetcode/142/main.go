/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
    // Solution using Floyd's Cycle Finding Alorithm
    slow := head
    fast := head
    for slow != nil && fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    // check for no loop
    if slow == nil || fast == nil || fast.Next == nil || slow != fast {
        return nil
    }
    // otherwise loop exists. find starting node by resetting slow to head and iterating both at same speed until they collide
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
