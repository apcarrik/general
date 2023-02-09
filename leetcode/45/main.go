// Notes:
//  - tried BFS, it used too much memory
//  - tried recursive search, it was too slow
//  - this solution uses DFS with memoization

type Node struct {
    Idx int
    PathLen int
    Next *Node
}

func jump(nums []int) int {
    // Idea: Use DFS with memoization to keep from re-computing unnessecarily

    n := len(nums)

    // Edge case
    if n==1 {
        return 0
    }

    // Create lookup table of shortest paths and initialize every element to +inf (n in this case, as largest possible path is n-1), except for first element which is 0
    shortestPaths := make([]int, n)
    shortestPaths[0] = 0
    for i:=0; i<n; i++ {
        shortestPaths[i] = n
    }

    sentinel := Node{-1,0,nil}
    head := &Node{0,0,&sentinel}

    for true { // max # of iterations is n!, which is too large for for loop iteration with integers
        if *head == sentinel {
            break
        }

        // pop head
        nodeIdx := head.Idx
        nodePathLen := head.PathLen
        nodeVal := nums[nodeIdx]
        head = head.Next

        for offset:=1; offset<=nodeVal; offset++{
            newIdx := nodeIdx+offset
            newPathLen := nodePathLen+1
            if newIdx < n {
                if newPathLen < shortestPaths[newIdx] {
                    shortestPaths[newIdx] = newPathLen
                    newNode := Node{newIdx,newPathLen,head}
                    head = &newNode
                }
            }
        }
    }
    return shortestPaths[n-1]

}
