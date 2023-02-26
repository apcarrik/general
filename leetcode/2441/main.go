import (
    "sort"
)

func findMaxK(nums []int) int {
    // Idea: sort array. Then, starting from the point where the negative and positive numbers meet, iterate through the negatives and positives simelaneously, starting from beginning and end of sorted array. Once you find a hit, it will be the largest pair.

    sort.Ints(nums)

    posPointer := len(nums)-1
    negPointer := 0

    for i:=0;i<len(nums);i++{
        pNum := nums[posPointer]
        nNum := nums[negPointer]
        nNumAbs := -1*nNum
        if pNum < 0 || nNum > 0 {
            break
        }
        if nNumAbs == pNum {
            return pNum
        } else if nNumAbs > pNum {
            negPointer++
        } else {
            posPointer--
        }
    }
    return -1

}
