// func mincostTickets(days []int, costs []int) int {
    
// }

func mincostTickets( days []int, costs []int) int {
	// create minCost array to hold min cost of day corresponding to days index
	minCost := []int{}
	for i:=0; i<len(days); i++{
		minCost = append(minCost, 365001)
	}
	
	// iterate through each day in days, starting from end and calculate its minimum cost
	for i:= len(days)-1; i>=0; i--{
		cost1 := costs[0]
		cost7 := costs[1]
		cost30 := costs[2]
		if i < len(days) - 1 {
			cost1 += minCost[i+1]
			cost7Added := false
			for j:=i; j<len(days); j++{
				if ! cost7Added && days[j] >= days[i]+7 {
					cost7 += minCost[j]
					cost7Added = true
				}
				if days[j] >= days[i]+30 {
					cost30 += minCost[j]
					break
				}
			}
		}
		//fmt.Println(days[i], cost1, cost7, cost30)
		switch{
		case cost1 <= cost7 && cost1 <= cost30:
				minCost[i] = cost1
		case cost7 <= cost1 && cost7 <= cost30:
				minCost[i] = cost7
		case cost30<= cost1 && cost30<= cost7:
				minCost[i] = cost30
		} 
	}
	    

	// return minCost for first day
	//fmt.Println(minCost)
	return minCost[0]

}
