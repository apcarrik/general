/*
   Notes:   - what if we found the "pools" of fixed-bound subarrays? i.e. largest fixed-bound subarrays in nums
            - we know that all subarrays under that would count, and there will be TriangularNumber(len(subarray))
            of them, where TriangularNumber(n) = (n^2+n)/2 (!!!ONLY if maxK = minK!!!)
*/
// import "fmt"
func triangularNumber(n int64) int64 {
    return (n*n+n)/2
}

func countSubarraysSameMinMax(nums []int, minK int) int64 {
    // Idea: iterate through each element of nums, until you find subarray with min==minK. Note the starting index, and keep iterating until you reach an element that is != minK. Then, append the triangular number of the length of this subarray and start over at the next element.

    fixedBoundCount := int64(0)
    startIdx := -1
    for i,num := range nums {
        if num == minK {
            // within bounds, keep track
            if startIdx == -1 {
                startIdx = i
            }
        } else {
            // out of bounds, add to fixedBoundCount if fixedBound found
            // fmt.Println(i, "startIdx",startIdx)

            if startIdx > -1  {
                fixedBoundCount += triangularNumber(int64(i-startIdx))
                // fmt.Println("fixedBoundCount",fixedBoundCount)
            }
            startIdx = -1
        }
    }
    if startIdx > -1  {
        fixedBoundCount += triangularNumber(int64(len(nums) - startIdx))
        // fmt.Println("fixedBoundCount",fixedBoundCount)
    }
    return fixedBoundCount
}

/*
   Brute-force solution
*/
func countSubarrays(nums []int, minK int, maxK int) int64 {
    // Idea: check all subarrays, keeping track of max/min for each and which are fixed-bound
    // - To do this, iterate through each index of nums as start of array, and iterate through remaining indicies of nums as end of array. This allows you to keep track of min/max for each start index efficiently.
    if minK == maxK {
        return countSubarraysSameMinMax(nums, minK)
    } else {
        fixedBound := 0
        for i:=0; i<len(nums); i++{
            max := 0
            min := 1000001
            jLoop:
            for j:=i; j<len(nums); j++{
                if nums[j] > max {
                    max = nums[j]
                } 
                if nums[j] < min {
                    min = nums[j]
                }
                if min < minK || max > maxK {
                    break jLoop
                }
                if min == minK && max == maxK {
                    fixedBound++
                }
            }
        }
        return int64(fixedBound) 
    }
}

/*
   Notes:   - the smallest fixed-bound subarray will have minimum and maximum as ends
            - maybe we could try to find these subarrays with min/max at ends that are fixed-bound, and then grow 
            them from there
            - Idea: find all indicies of minK and maxK. For each index of minK, search through all indicies of maxK 
            until you find a fixed-bound subarray between them. Then, try to grow subarray in each direction until 
            it no longer satifsies fixed-bound subarray, or you hit another minK or maxK.
*/
