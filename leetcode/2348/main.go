/* Notes:
    - each length subarray has a deterministic number of smaller subarrays (so we only need to count max length sub arrays)
    - idea: go through nums, searching for the longest subarrays present. Then add in the smaller subarrays based on those.
    - can store subarrays in map with key: length of subarray, value: number of occurences in nums
    - for subarray of length n, the sum of all subarrays that can be made of it is the nth triangular number: (n*(n+1))/2
    [0] = [0]: 1
    [0,0] = [0],[0],[0,0] : 1 + 2 = 3
    [0,0,0] = [0,0,0],[0,0],[0,0],[0],[0],[0] : 1 + 2 + 3 = 6
    [0,0,0,0] = 1 + 2 + 3 + 4 = 10

*/

func zeroFilledSubarray(nums []int) int64 {
    subs := map [int]int{}

    // count number of largest length subarrays found in nums
    startIdx := -1
    for i,num := range nums {
        if num == 0 && startIdx == -1 {
            startIdx = i
        } else if num != 0 && startIdx != -1 {
            subLen := i-startIdx
            // add to subs
            if _,ok := subs[subLen]; ok {
                subs[subLen]++
            } else {
                subs[subLen] = 1
            }
            startIdx = -1
        }
    }

    // check if last index is subarray
    if startIdx != -1 {
        subLen := len(nums)-startIdx
        // add to subs
        if _,ok := subs[subLen]; ok {
            subs[subLen]++
        } else {
            subs[subLen] = 1
        }
        startIdx = -1
    }

    // calculate additional subarrays based on length of longest subarray
    retVal := int64(0)
    for length, count := range subs {
        retVal += (int64(length)*(int64(length)+1))/2 * int64(count)
    }
    return retVal


}
