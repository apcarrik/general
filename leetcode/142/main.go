/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /* Idea: keep a set of nodes seen. Once you see a node again, return that node

 */
func detectCycle(head *ListNode) *ListNode {
    nodeSet := map [*ListNode]bool{}
    for iter := head; iter != nil; iter = iter.Next {
        if _,ok := nodeSet[iter]; ok {
            return iter
        }
        nodeSet[iter] = true
    }
    return nil    
}
