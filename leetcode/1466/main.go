/* Notes:
    - Brute force approach: starting from node 0, search for all nodes that point to this node (or point to it when turned around)
        - O(n^2) (look through n connetions for each of n nodes)
    - could we use a map? what's the key? first node?
    - Idea: use two maps, one for forward edges, one for backward edges
        - Each edge would be duplicated, one in each map. once we traversed that edge, we'd need to remove the other one.
        - go through entire tree, counting the number of forward edges used

    ex n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
    forward: [  0:[1]
                1:[3]
                x2:[x3]
                4:[x0, 5]
             ]
    backward: [  0:[4]
                x1:[x0]
                3:[x1,2]
                x5:[x4]
             ]


*/

type Node struct {
    Val int
    Next *Node
}

func minReorder(n int, connections [][]int) int {
    // Make forward and backward maps
    forward := map[int][]int{}
    backward := map[int][]int{}
    for _,conn := range connections {
        // add to forward
        if _,ok := forward[conn[0]]; ok {
            forward[conn[0]] = append(forward[conn[0]], conn[1])
        } else {
            forward[conn[0]] = []int{conn[1]}
        }
        // add to backward
        if _,ok := backward[conn[1]]; ok {
            backward[conn[1]] = append(backward[conn[1]], conn[0])
        } else {
            backward[conn[1]] = []int{conn[0]}
        }
    }

    // Starting from 0, traverse through tree, removing incorrectly aligned nodes
    head := &Node{Val:0}
    tail := head
    for head != nil {
        // check forward
        if _,ok := forward[head.Val]; ok {
            src := head.Val
            for _,dest := range forward[head.Val]{
                src2 := dest
                for i,dest2 := range backward[src2] {
                    if dest2 == src {
                        // remove from backward
                        if i == len(backward[src2])-1 {
                            backward[src2] = backward[src2][:i]
                        } else {
                            backward[src2] = append(backward[src2][:i], backward[src2][i+1:]...)
                        }
                    }
                }
                newTail := &Node{Val:dest}
                tail.Next = newTail
                tail = tail.Next
            }
        }
        
        // check backward
        if _,ok := backward[head.Val]; ok {
            src := head.Val
            for _,dest := range backward[head.Val]{
                src2 := dest
                for i,dest2 := range forward[src2] {
                    if dest2 == src {
                        // remove from forward
                        if i == len(forward[src2])-1 {
                            forward[src2] = forward[src2][:i]
                        } else {
                            forward[src2] = append(forward[src2][:i], forward[src2][i+1:]...)
                        }
                    }
                }
                newTail := &Node{Val:dest}
                tail.Next = newTail
                tail = tail.Next
            }
        }
        
        // move to next in queue
        head = head.Next
    }

    // Now, count the number of values in forward
    reversed := 0
    for _,v := range forward {
        reversed += len(v)
    }
    return reversed
    
}
