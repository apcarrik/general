/*Notes:
    - Binary search problem
    - max k: maximum pile of bananas
    - min k: 1
    - note: low may be +1 or -1 of actual answer, or next highest stack. Hence edge cases check at end.

*/

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

    // find max pile for high
    maxPile := 0
    for _,pile := range piles {
        if pile > maxPile {
            maxPile = pile
        }
    }

    // binary search
    low := 1
    high := maxPile
    for high > low {
        mid := (high-low)/2 + low
        hours := getHours(&piles, mid)
        if hours > h {
            low = mid + 1            
        } else {
            high = mid - 1
        }
    }

    // edge cases
    if low == maxPile {
        secondMaxPile := 0
        for _,pile := range piles {
            if pile > secondMaxPile && pile < maxPile {
                secondMaxPile = pile
            }
        }
        if secondMaxPile > 0 && getHours(&piles, secondMaxPile) <= h {
            return secondMaxPile
        }
    }
    
    if low-1 > 0 && getHours(&piles, low-1) <= h {
        return low-1
    } 
    
    if getHours(&piles, low) > h {
        return low+1
    }
    
    return low
    
}
