func updateNodes(len int, cycleLenPtr *[]int, nodeMapPtr *map[int]int) {
	for k,_ := range *nodeMapPtr {
		(*cycleLenPtr)[k] = len
	}
	return
}

func longestCycle(edges []int) int {
	cycleLen := []int{}

	// initialize cycleLen to -1
	for i:=0; i<len(edges); i++ {
		cycleLen = append(cycleLen, -1)	
  }	

	// iterate through edges
	for _,node := range edges {
		// find cycle length of this edge, and which element starts cycle
		nodeMap := map[int]int{} // k:nodeID, v: iteration seen
		for iter:=0; iter<=len(edges); iter++ {
			if node == -1 {
				// optimization: update everything in this nodes path
				updateNodes(0, &cycleLen, &nodeMap)
				break
			}
			if cycleLen[node] != -1 {
				// optimization: update everything in this nodes path
				updateNodes(cycleLen[node], &cycleLen, &nodeMap)			
				break
			}
			if _,ok := nodeMap[node]; ok {
				// found loop
				// optimization: update everything in this nodes path
				updateNodes(iter - nodeMap[node], &cycleLen, &nodeMap)
				break
			} else {
					nodeMap[node] = iter	
			}			
			node = edges[node]
		}
	}

	// find and return max of cycleLen
	maxLen := 0
	for _,len := range cycleLen {
		if len > maxLen {
			maxLen = len
		}
	}
	if maxLen == 0 {
		maxLen = -1
	}
	return maxLen
    
}
