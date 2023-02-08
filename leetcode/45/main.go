func jump(nums []int) int {
    // Idea: Consider nums an encoding of a graph. Use Dijkstra's algorithm (or similar BFS-based algorithm) to find shortest path between nodes of first element to last element of nums. Starting at the first.
    // BFS search uses a Priority Queue to keep nodes in order to search in correct order
    n := len(nums)

    // Keep slice denoting shortest path to each node from node 0.
    shortestPath := make([]int, n)
    // initialize shortest path of all nodes to +inf
    for i:=1;i<n;i++{ // first node has path of 0
        shortestPath[i] = n+1
    }


    // I'm going to use a slice to implement Priority Queue
    // two queues: nodeQ indicates the index of node in nums, pathQ indicates current path of node on nodeQ
    nodeQ := []int{0}
    pathQ := []int{0}
    for w:= 0; w<n*n; w++ { // max # of iterations is n^2
        if len(nodeQ) == 0 {
            break
        }

        // Pop priority from nodeQ and pathQ
        thisNode := nums[nodeQ[0]]
        nodeQ := nodeQ[1:]
        thisPathLen := pathQ[0]
        pathQ := pathQ[1:]

        // TODO: implement Dijkstra's algorithm


    }
    return shortestPath[n-1]
}
