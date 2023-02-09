import "fmt"
type QNode struct {
    Value int
    Next *QNode
}

func jump(nums []int) int {
    // Idea: Consider nums an encoding of a graph. Use Dijkstra's algorithm (or similar BFS-based algorithm) to find shortest path between nodes of first element to last element of nums. Starting at the first.
    // BFS search uses a Queue to keep nodes in order to search in correct order
    n := len(nums)

    // Keep slice denoting shortest path to each node from node 0 (indicies correspond to nums)
    shortestPath := make([]int, n)
    // initialize shortest path of all nodes to +inf
    shortestPath[0] = 0 // first node has path of 0
    for i:=1;i<n;i++{ 
        shortestPath[i] = n+1
    }

    
    // two queues: nodeQ indicates the index of node in nums, pathQ indicates current path of node on nodeQ
    sentinel := QNode{-1,nil}
    nodeQHead := QNode{0,&sentinel}
    nodeQTail := nodeQHead
    pathQHead := QNode{0,&sentinel}
    pathQTail := pathQHead
    
    for true { // Warning: max # of iterations is n!, but I couldn't get for loop to bound that high
        if nodeQHead == sentinel {
            break
        }
        // Pop priority from nodeQ and pathQ
        thisNodeIdx := nodeQHead.Value
        nodeQHead = *nodeQHead.Next
        thisPathLen := pathQHead.Value
        pathQHead = *pathQHead.Next

        // Check if
        if thisPathLen <= shortestPath[thisNodeIdx] { 
            shortestPath[thisNodeIdx] = thisPathLen
            if thisNodeIdx < (n-1) {
                // Add all nodes we can reach from this one to queue
                for i:=1; i<=nums[thisNodeIdx]; i++{
                    if thisNodeIdx+i<n {
                        newNode := QNode{thisNodeIdx+i,&sentinel}
                        if nodeQHead == sentinel {
                            nodeQHead = newNode
                            nodeQTail = nodeQHead
                        } else {
                            nodeQTail.Next = &newNode
                            nodeQTail = newNode
                        }
                        // nodeQ = append(nodeQ, thisNodeIdx+i)
                        newNode = QNode{thisPathLen+1,&sentinel}
                        if pathQHead == sentinel {
                            pathQHead = newNode
                            pathQTail = pathQHead
                        } else {
                            pathQTail.Next = &newNode
                            pathQTail = newNode
                        }
                        // pathQ = append(pathQ, thisPathLen+1)
                    }
                }
            }
        }
    }

    return shortestPath[n-1]
}
