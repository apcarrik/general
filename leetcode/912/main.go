// import "fmt"
func mergeSort(numsPtr *[]int, min, max int) {
    nums := *numsPtr
    // Find midpoint
    mid := (max-min)/2+min

    // Stop dividing and sort when max is < 3 away from min
    if max-min < 1 {
        switch max-min {   
        case 1:
            if nums[max] < nums[min] {
                (*numsPtr)[min], (*numsPtr)[max] = (*numsPtr)[max], (*numsPtr)[min]
            }
            return
        case 0:
            return
        }
    }


    // Divide array in half and recusivley call mergeSort on two halves
    h1Min := min
    h1Max := mid
    h2Min := mid+1
    h2Max := max
    //fmt.Println("h1Min:", h1Min, "h1Max:", h1Max, "h2Min:",h2Min, "h2Max:",h2Max)
    mergeSort(numsPtr, h1Min, h1Max)
    mergeSort(numsPtr, h2Min, h2Max)

    // Order both sorted halves together in new array
    newArr := []int{}
    for h1Min <= h1Max || h2Min <= h2Max {
        if h1Min > h1Max {
            newArr = append(newArr,nums[h2Min])
            h2Min++
            continue
        }
        if h2Min > h2Max {
            newArr = append(newArr,nums[h1Min])
            h1Min++
            continue
        }

        if nums[h1Min] < nums[h2Min] {
            newArr = append(newArr,nums[h1Min])
            h1Min++
        } else {
            newArr = append(newArr,nums[h2Min])
            h2Min++
        }
    }

    // Copy elements of new array into nums between min and max
    for i := 0; i <= len(newArr) - 1; i++{
        (*numsPtr)[min+i] = newArr[i]
    }
    //fmt.Println("h1Min:", h1Min, "h1Max:", h1Max, "h2Min:",h2Min, "h2Max:",h2Max, "newArr", newArr, "nums", nums)
    return

}

func sortArray(nums []int) []int {
    // Thoughts: O(n log(n)) complexity means some type of binary search applied (i.e. quicksort, mergesort)
    // Idea: Implement mergesort (I think... the one that divides the search area in 2 subproblems until you only  have 2 elements to compare, then merges the results)

    mergeSort(&nums, 0, len(nums)-1)
    return nums

}
