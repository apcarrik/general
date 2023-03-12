/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* Idea: continually add the minimum element of all the linked lists to the new list.

*/


func mergeKLists(lists []*ListNode) *ListNode {

    newHead := &ListNode{ // temporary sentinel node
        Val: -10001,
    }
    newNode := newHead
    for true {
        min := 10001
        minI := -1
        for i,list := range lists {
            if list != nil && list.Val < min {
                min = list.Val
                minI = i
            }
        }
        if minI == -1 {
            break
        }
        newNode.Next = &ListNode{ // temporary sentinel node
            Val: min,
        }
        newNode = newNode.Next
        lists[minI] = lists[minI].Next
    }

    return newHead.Next // ignore sentinel node
    
}
