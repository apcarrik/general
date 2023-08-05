/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type ExtraNode struct {
	N *Node
	NeighborSet map[int]bool
}

func createNode(node *Node, nodeMapPtr *map[int]*ExtraNode) {
	nodeMap := *nodeMapPtr
	// check if node is in map, create if not
	if _,ok := nodeMap[node.Val]; !ok {
		nodeMap[node.Val] = &ExtraNode{N:&Node{Val:node.Val}, NeighborSet:map[int]bool{}}
	}
	clone := nodeMap[node.Val].N
	cloneSet := nodeMap[node.Val].NeighborSet 
	// check each neighbor of existing node
	for _,neighbor := range node.Neighbors {
		if _,ok := nodeMap[neighbor.Val]; ok {
			// check if neighbor has this node in its neighbors
			neighborClone := nodeMap[neighbor.Val].N
			neighborCloneSet := nodeMap[neighbor.Val].NeighborSet
            if _,ok := neighborCloneSet[clone.Val]; !ok {
				neighborClone.Neighbors = append(neighborClone.Neighbors, clone)
				neighborCloneSet[clone.Val] = true
			}
		} else {
			// recursively call find Neighbors before we add to our neighbors
			createNode(neighbor, nodeMapPtr)			
		}
		// check if neighbor node is in our neighbors and if not add it
		neighborClone := nodeMap[neighbor.Val].N
		if _,ok := cloneSet[neighbor.Val]; !ok {
			clone.Neighbors = append(clone.Neighbors, neighborClone)
			cloneSet[neighbor.Val] = true
		}
	}

}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
  nodeMap := map[int]*ExtraNode{}
	createNode(node, &nodeMap)
	return nodeMap[node.Val].N
}
