/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Idea: keep two pointers, one n away from the first. Iterate each through list until
    // you find the end, and then the other pointer will have the nth node from the end.
    end := head
    for i:=1; i<n; i++{
        end = end.Next
    }
    nthFromEnd := head
    var nthFromEndParent *ListNode
    for i:=0; i<30; i++{
        if end.Next == nil {
            break
        }
        nthFromEndParent = nthFromEnd
        nthFromEnd = nthFromEnd.Next
        end = end.Next
    }
    if nthFromEndParent == nil {
        return nthFromEnd.Next
    }
    nthFromEndParent.Next = nthFromEnd.Next    
    return head
}
