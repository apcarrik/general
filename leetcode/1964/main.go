/* Notes:
    - for each element of obstacles, we want the number of obstacles before this one that are greater than or equal to all obstacles before it.
    - I think we can start from beginning and compute obstacle course length for each element
    - We can use pre-computed obstacle course lengths to save time, by searching one index to the left of current until you hit one that is less than or equal to current and summing that with current
*/

/* Problem: this won't scale well. Worst case every element is strictly decreasing, it will be a pyramid sum, so O(n^2) runtime
    - Idea: keep a map with obstacle lengths as keys, and array of indicies for that obstacle length as value, with larger indicies at the front.

    Idea: keep rolling map with key = longest values and value = smallest height in obstacle course. You'll need to keep sorting the keys each iteration of for loop around obstacles. In each iteration, check each longest value if the max height < this height, else move to the next one

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
    // iterate through obstacles, recording the longest obstacle course using obstacleLookup
    longest := []int{}
    longestLookup := map [int]int{} // key: longeset obstacle course, value: smallest height to have this obstacle course // OLD: value: largest height in obstacle course
    longestLookupKeys := []int{} // ordered list of keys from longest lookup (smallest to largest)
    for _, height := range obstacles {

        // find largest previous obstacle course
        prevSum := 0
        for i:=len(longestLookupKeys) - 1; i>=0; i-- {
            k := longestLookupKeys[i]
            if longestLookup[k] <= height {
                prevSum = k
                break                
            }
        }
        // update longest
        newSum := prevSum +1
        longest = append(longest, newSum)

        if _,ok := longestLookup[newSum]; ok {
            if longestLookup[newSum] > height {
                longestLookup[newSum] = height
            }
        } else {
            longestLookup[newSum] = height
            // add newSum to longest lookup keys
            newPosition := binarySearch(newSum, &longestLookupKeys)
            if newPosition == -1 {
                longestLookupKeys = append([]int{newSum}, longestLookupKeys...)
            } else if newPosition < len(longestLookupKeys) - 1{
                pre := longestLookupKeys[:newPosition+1]
                post := longestLookupKeys[newPosition+1:]
                longestLookupKeys = append(append(pre, newSum), post...)
            } else {
                longestLookupKeys = append(longestLookupKeys, newSum)
            }
        }
    }
    return longest
    
}
