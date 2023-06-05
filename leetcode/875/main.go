func getHours(pilesPtr *[]int, k int) int {
    piles := *pilesPtr
    hours := 0
    for _,pile := range piles {
        hours += pile / k
        if pile % k != 0 {
            hours++
        }
    }
    return hours
}

func minEatingSpeed(piles []int, h int) int {
    low := 1
    high := 1000000001
    for high > low {
        mid := (high-low)/2 + low
        hours := getHours(&piles, mid)
        if hours > h {
            low = mid + 1            
        } else {
            high = mid - 1
        }
    }    
    if getHours(&piles, low) > h {
        return low+1
    }    
    return low    
}
