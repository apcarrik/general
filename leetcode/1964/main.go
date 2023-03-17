/* Notes:
    - for each element of obstacles, we want the number of obstacles before this one that are greater than or equal to all obstacles before it.
    - I think we can start from beginning and compute obstacle course length for each element
    - We can use pre-computed obstacle course lengths to save time, a form of memoization
*/

/* Solution: Iterate through each index (height) of obstacles from beginning to end. Keep track of the longest obstacle course at each height (longest), and keep a lookup table (longestLookup) with {key: obstacle course length, value: smallest height to have this obstacle course length} and a separate slice with the keys from this lookup table (longestLookupKeys) that is kept in increasing order. 

For each height in obstacles, find if any previous obstacle course lengths have a minimum height less than or equal to this height, and if so use that length + 1 as the length of this obstacle course. Add this new obstacle course length to longest, and update longestLookup and longestLookupKeys as appropriate.

*/

func binarySearch(x int, thingsPtr *[]int) int { // gaurentees to return index of closest element from things that is <= x
    things := *thingsPtr
    if len(things) == 0 {
        return -1
    }
    low := 0
    high := len(things)
    for low < high {
        mid := (high - low) / 2 + low
        if things[mid] == x {
            return mid
        } else if things[mid] > x {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    if low >= len(things)-1 || things[low] > x {
        low--
    }     
    return low

}


func longestObstacleCourseAtEachPosition(obstacles []int) []int {
    longest := []int{}
    longestLookup := map [int]int{} // key: obstacle course length, value: smallest height to have this obstacle course
    longestLookupKeys := []int{} // ordered list of keys from longest lookup (smallest to largest)

    // iterate through obstacles, recording the longest obstacle course using obstacleLookup
    for _, height := range obstacles {

        // find largest previous obstacle course
        prevLength := 0
        for i:=len(longestLookupKeys) - 1; i>=0; i-- {
            k := longestLookupKeys[i]
            if longestLookup[k] <= height {
                prevLength = k
                break                
            }
        }
        // update longest
        newLength := prevLength +1
        longest = append(longest, newLength)

        // update longestLookup
        if _,ok := longestLookup[newLength]; ok {
            if longestLookup[newLength] > height {
                longestLookup[newLength] = height
            }
        } else {
            longestLookup[newLength] = height
            // add newLength to longestLookupKeys
            newPosition := binarySearch(newLength, &longestLookupKeys) // keep longestLookupKeys in increasing order
            if newPosition == -1 {
                longestLookupKeys = append([]int{newLength}, longestLookupKeys...)
            } else if newPosition < len(longestLookupKeys) - 1{
                pre := longestLookupKeys[:newPosition+1]
                post := longestLookupKeys[newPosition+1:]
                longestLookupKeys = append(append(pre, newLength), post...)
            } else {
                longestLookupKeys = append(longestLookupKeys, newLength)
            }
        }
    }
    return longest
    
}
