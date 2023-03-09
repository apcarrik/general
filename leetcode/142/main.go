/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
    // this function implements detectCycle with only a constant amount of memory
    // Idea: Send out two 'runners' at different speeds to traverse through the linked list. Once the front 'runner' intercepts a slower one, you know there is a loop. Keep adding an offset to the second runner until you catch where the first runner points to the second but the second had a parent that is different than the first.

    // edge case: head is nil
    if head == nil {
        return nil
    }

    // edge case: 1 node with cycle
    if head.Next == head {
        return head
    }

    for r2Offset := 0; r2Offset < 10001; r2Offset++{
        r1 := head
        r1Steps := 0
        r2 := head
        r2Parent := r2 
        for r1.Next != r2{
            // move r1
            r1 = r1.Next
            r1Steps++

            // check if r1 is now nil, indicating no loop
            if r1 == nil {
                return nil
            }

            // move r2 if applicable
            if (r1Steps-r2Offset) % 2 == 1 && r1Steps > 0 {
                r2Parent = r2
                r2 = r2.Next
            }
        }
        // check if r1 is right behind r2
        if r1.Next == r2 {
            // check if r2's parent is different from r1 and return. else, incrase r2offset and start over
            if r2Parent != r1 {
                return r2
            }
            continue
        }
    }

    return nil
}
