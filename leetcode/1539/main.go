func findKthPositive(arr []int, k int) int {
    // edge case
    if k<arr[0] {
        return k
    }

    missingArrElementsCount :=  arr[len(arr)-1] - len(arr)
    if k <= missingArrElementsCount { // perform binary search for missing element in array
        min := 0
        max := len(arr)-1
        for min<max {
            mid := (max-min) / 2 + min
            missingArrElementsCount = arr[mid] - mid - 1
            switch {
                case k > missingArrElementsCount:
                    min = mid+1
                case k <= missingArrElementsCount:
                    max = mid
            }
        }
        min--
        missingArrElementsCount = arr[min] - min
        for i:=arr[min]+1; i<arr[max]; i++ {
            if missingArrElementsCount == k {
                return i
            }
            missingArrElementsCount++
        }

    }
    // missing element is after array
    return arr[len(arr)-1] + (k-missingArrElementsCount)
}
