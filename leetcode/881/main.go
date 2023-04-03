// func numRescueBoats(people []int, limit int) int {
    
// }


import (
	"sort"
)

func binSearch(needle int, haystack []int) int {
	// return index of element <= needle in haystack, or -1 if it doesn't exist
	//fmt.Println(needle, haystack)
	low := 0 
	high := len(haystack)-1
	for high > low {
		mid := (high - low) / 2 + low
		if haystack[mid] <= needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	//fmt.Println("low",low)
	if low >= len(haystack)-1 {
    if haystack[len(haystack)-1] <= needle {
			return len(haystack)-1
		} else {
			if low > 0 && haystack[low-1] <= needle {
				return low-1
			}
			return -1
		}
	}
	if low == 0 {
		if haystack[0] > needle {
			return -1
		} else if haystack[1] <= needle {
			return 1
		}
	}
	if haystack[low] > needle{
		low--
	}
	return low
}

func numRescueBoats(people []int, limit int) int {
	// sort people
	sort.Ints(people)
	//fmt.Println("sorted:", people)
	// iterate through people from heaviest to lightest
	boats := 0
	for i := len(people)-1; i >= 0; i--{
		boats++
		first := people[i]
		people = people[:len(people)-1] 
		if len(people) > 0 {
			capacity := limit - first
			second := binSearch(capacity, people)
			//fmt.Println(second)
			if second != -1 {
				if second == 0 {
					if len(people) == 1 {
						break
					}
					people = people[1:]
				} else if second == len(people)-1 {
					people = people[:len(people)-1]
        } else {
					people = append(people[:second], people[second+1:]...)
        }
				i--
			}
		}
		//fmt.Println("boats:", boats, "people:", people, "i:", i)
	}


	return boats
	
}
