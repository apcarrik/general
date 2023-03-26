func longestCycle(edges []int) int {
	cycleLen := []int{}

	// initialize cycleLen to -1
	for i:=0; i<len(edges); i++ {
		cycleLen = append(cycleLen, -1)	
  }	

	// iterate through edges
	for i,node := range edges {
		// find cycle length of this edge, and which element starts cycle
		nodeMap := map[int]int{} // k:nodeID, v: iteration seen
		for iter:=0; iter<=len(edges); iter++ {
			if node == -1 {
				cycleLen[i] = 0
				break
			}
			if cycleLen[node] != -1 {
				cycleLen[i] = cycleLen[node]				
				break
			}
			if _,ok := nodeMap[node]; ok {
				// found loop
				cycleLen[node] = iter - nodeMap[node]
				cycleLen[i] = cycleLen[node]
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
