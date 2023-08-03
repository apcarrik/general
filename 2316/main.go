// func countPairs(n int, edges [][]int) int64 {
    
// }


type CCNode struct {
	Nodes []int
}

func countPairs(n int, edges [][]int) int64 {
	nMap := map[int]*CCNode{}
	for i:=0; i<n; i++ {
		nMap[i] = &CCNode{Nodes: []int{i}}
	}
	
	for _,edge := range edges {
		sm := nMap[edge[0]]
		lg := nMap[edge[1]]
		if sm == lg {
			continue
		}
		if len(sm.Nodes) > len(lg.Nodes) {
			sm, lg = lg, sm			
		}
		//fmt.Println(edge[0], sm, edge[1], lg)
		for _,node := range sm.Nodes {
			nMap[node] = lg
			//fmt.Println(node, nMap[node])
		}
		lg.Nodes = append(lg.Nodes, sm.Nodes...)
		//fmt.Println(nMap)
		// for i,n := range nMap {
			//fmt.Println(i,n)
		// }
		//fmt.Println(nMap)
	}
	//fmt.Println(nMap[0])
	ccs := map [*CCNode]bool{} // Set of Connected Components
	nodesSeen := 0
	retVal := int64(0)
	for _,cc := range nMap {
		if _,ok := ccs[cc]; !ok {
			ccs[cc] = true
			ccNodeCount := len(cc.Nodes)
			retVal += int64((n - ccNodeCount - nodesSeen) * ccNodeCount)
			nodesSeen += ccNodeCount
		}
  } 

    return retVal
}
