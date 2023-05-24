func numOfSubarrays(arr []int) int {
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

    numOddSums := 0
    // iterate through each element of arr, adding up number of odd subarray sums for each element
    for idx,elem := range arr {
        if numOdd %2 == 0 && numOdd > 0 { // group 1
            numOddSums += group1[idx]
        } else if numOdd %2 == 1 { // group 2
            numOddSums += (n-idx) - group1[idx]
        }
        numOdd -= elem % 2
    }

    return numOddSums % (1000000007)
    
}
