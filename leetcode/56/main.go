/* Notes:
    - it would probably help to sort all intervals by start
    - for each interval, search all other intervals for overlapping ones. You will only need to search intervals with start less than or equal to current intervals end


*/

import (
    "fmt"
)

type intervalsSorted [][]int

func (a intervalsSorted) Len() int { return len(a) }
func (a intervalsSorted) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a intervalsSorted) Less(i, j int) bool { return a[i][0] < a[j][0]}

func merge(intervals [][]int) [][]int {

    // sort all intervals by start
    sort.Sort(intervalsSorted(intervals))

    // search for overlapping intervals
    for i:=0; i<len(intervals)-1; i++ {
        current := intervals[i]
        compare := intervals[i+1]
        if compare[0] > current[1] {
            continue
        } else {
            // merge
            newInterval := []int{current[0], compare[1]}
            if current[0] > compare[0] {
                newInterval[0] = compare[0]
            }
            if current[1] > compare[1] {
                newInterval[1] = current[1]
            }
            // update intervals
            if i == len(intervals)-2 {
                intervals = append(intervals[:i], newInterval)
            } else {
                intervals = append(append(intervals[:i], newInterval), intervals[i+2:]...)
            }
            i--
        }        

    }

    return intervals
}
