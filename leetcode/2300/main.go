// func successfulPairs(spells []int, potions []int, success int64) []int {
    
// }

import (
	"sort"
)

func binSearch(needle int64, haystack []int) int {
	// search for needle in haystack and return -1 if not found, else the smallest element of haystack that is greater than or equal to needle
	low := 0
	high := len(haystack)


	for high > low {
		mid := (high-low)/2 + low
		if int64(haystack[mid]) >= needle {
			high = mid -1
    } else {
			low = mid + 1
		}
	
	}
	//fmt.Println("needle:", needle, "low:", low)
	if low >= len(haystack){
		return -1
	}
	if low <= len(haystack)-1 && int64(haystack[low]) < needle  {
		low++ 
	}
  if int64(haystack[low]) < needle {
		return -1
	}
	
	return low
}

func successfulPairs(spells []int, potions[]int, success int64) []int {

	// Sort potions
	sort.Ints(potions)		
		
	// iterate throug spells, perofrming binary search for potions within bounds and updating spell with number of potions that work for it
	for si,spell := range spells {
		extra := int64(0)
		if success % int64(spell) != 0 {
			extra = 1
		}
		lim := success / int64(spell) + extra
		minPotionIdx := binSearch(lim, potions)
		//fmt.Println(minPotionIdx)
		if minPotionIdx != -1 && minPotionIdx < len(potions) {
			spells[si] = len(potions) - minPotionIdx
		} else {
			spells[si] = 0
		}

	}

	return spells


}
