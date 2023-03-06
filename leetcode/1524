/* Notes
    - odd sum can only happen when there is an odd number of odds in summands
    - Brute force: find all subarrays' sums and count the number of odd, O(n^2) time complexity
    - Array is not gaurenteed to be sorted, but it shouldn't need to be sorted
    - Idea: Iterate through each index of the array, adding up the number of subarrays from that index until end 
    that have an odd sum. 
        - Do this efficently by making a lookup table for each index, corresponding to two groups based on if # odd
         from idx to end is even, or if it is odd.

*/

func bruteForce(arr []int) int { // O(n^2) solution
    n := len(arr)
    numOddSums := 0
    // for each index, find number of odd sums 
    for start,_ := range arr {
        numOdds:=0
        for end := start; end < n; end++{
            elem := arr[end]
            numOdds += elem % 2
            if numOdds % 2 == 1 {
                numOddSums += 1
            }
        }
    }

    return numOddSums % (1000000007)

}


func numOfSubarrays(arr []int) int {

    // return bruteForce(arr) // too slow

    // // Edge case: arr is len 1
    // if len(arr) == 1 {
    //     return arr[0] % 2
    // }

    n := len(arr)
    // find number of odd elements
    numOdd := 0
    for _,elem := range arr {
        numOdd += elem % 2
    }

    // mark elements belonging to group 1
    group1 := []int{}
    localNumOdd := numOdd
    for _,elem := range arr {
        if localNumOdd % 2 == 0  {
            if elem %2 == 0{
                group1 = append(group1, 0)
            } else {
                group1 = append(group1, 1)
            }
        } else {
            if elem %2 == 0{
                group1 = append(group1, 1)
            } else {
                group1 = append(group1, 0)
            }
        }
        localNumOdd -= elem % 2
    }

    // make group1 into lookup table
    for i:=len(group1)-2; i>=0; i--{
        group1[i] = group1[i] + group1[i+1]
    }
    //fmt.Println("group1:", group1, "numOdd", numOdd)

    numOddSums := 0
    // TODO: iterate through each element of arr, adding up number of odd subarray sums for each element
    for idx,elem := range arr {
        if numOdd %2 == 0 && numOdd > 0 { // group 1
            numOddSums += group1[idx]
            //fmt.Println("+",group1[idx],"group1")
        } else if numOdd %2 == 1 { // group 2
            numOddSums += (n-idx) - group1[idx]
            //fmt.Println("+",(n-idx) - group1[idx], "group2")
        }
        numOdd -= elem % 2
    }

    return numOddSums % (1000000007)
    
}
