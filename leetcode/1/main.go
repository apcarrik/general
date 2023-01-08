import (
    "sort"
    "fmt"
)

func binarySearch(needle int, haystack []int) int{
    if len(haystack) < 4 {
        for i:=0; i<len(haystack); i++ {
            if haystack[i] == needle{
                return i
            }
        }
        return -1
    }
    low := 0
    high := len(haystack) - 1

    for low <= high {
        median := (low + high) / 2
        // fmt.Println("   median: ", median)

        if haystack[median] < needle {
            low = median + 1
        } else {
            high = median - 1
        }
    }

    if low == len(haystack) || haystack[low] != needle {
        return -1
    }

    return low
}


func twoSum(nums []int, target int) []int {
    // 0. If length of nums is 2, return 0,1
    if len(nums) == 2{
        return []int{0,1}
    }

    // 1. Copy nums to nums2
    nums2 := make([]int, len(nums))
    copy(nums2,nums)

    // 2. Sort nums2
    sort.Ints(nums2)

    // 3. Create map of indecies from nums2 to nums
    nums_map := make([]int, len(nums))
    for i:=0; i<len(nums); i++{
        nums_map[i] = binarySearch(nums2[i],nums)
        fmt.Println("nums_map: ", nums_map)
        if i > 0 && nums_map[i-1] == nums_map[i]{
            nums_map[i] = binarySearch(nums2[i],nums[nums_map[i-1]+1:])+(nums_map[i-1]+1)
            fmt.Println(">nums_map: ", nums_map)
        }

    }

    // 4. Iterate over nums from smallest to largest, keeping curr as current number
    for i1 := 0; i1 < len(nums2)-1; i1++ {
        curr := nums2[i1]

        // 3. Perform binary search for index of target-curr
        i2:= binarySearch(target-curr, nums2[i1+1:])

        if i2 > -1{
            return []int{nums_map[i1], nums_map[i2+i1+1]}
        }

    }

    return []int{0, 0}
}
