/* Notes:
    - The minimum number of cables to fully connect n computers is n-1
    - What we're looking for the the number of connected components in the graph. 
    - If there are c connected components, then there will need to be c-1 cables moved to connect all of them
    Idea: ensure there are enough connections (>= n-1). If so, find the number of connected components, and return that number minus 1. Do this by keeping a slice with index corresponding to computer index and value being the connected component that computer is a part of.

*/

type ConnComp struct {
    Nodes []int // computers in connected component
}

func moveNodes(from, to *ConnComp, CCsPtr *[]*ConnComp) {
    CCs := *CCsPtr
    for _,node := range from.Nodes {
        to.Nodes = append(to.Nodes, node)
        CCs[node] = to
    }
    return
}

func makeConnected(n int, connections [][]int) int {
    // check if there exist enough connections
    if len(connections) < n-1 {
        return -1
    }

    // now, count the number of connected components

    // initialize slice of connected components, with one per computer
    CCs := []*ConnComp{}
    for i:=0; i<n; i++{
        CCs = append(CCs, &ConnComp{
            Nodes: []int{i},
        })
    }

    // Go through connections, updating connected components if necessary
    CCCount := n
    for _,connection := range connections {
        CC1 := CCs[connection[0]]
        CC2 := CCs[connection[1]]
        if CC1 != CC2 {
            // check which is bigger
            if len(CC1.Nodes) < len(CC2.Nodes) {
                // cc2 is bigger
                moveNodes(CC1, CC2, &CCs)
                CCs[connection[0]] = CC2
            } else {
                // cc1 is bigger
                moveNodes(CC2, CC1, &CCs)
                CCs[connection[1]] = CC1
            }
            CCCount--
            //fmt.Println("Coalescing nodes, connection:", connection, "CCCount:", CCCount)
        }
        //fmt.Println("CCs:",CCs)
    }
    return CCCount - 1
    
}
