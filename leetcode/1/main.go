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
        // fmt.Println("low:",low,"high",high,"haystack",haystack,"needle",needle)
        return -1
    }

    return low
}

func contains(needle int, haystack []int) int {
	for _, v := range haystack {
		if v == needle {
			return v
		}
	}
	return -1
}

func containsIndex(needle int, haystack []int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
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
    // fmt.Println("nums2: ", nums2)

    // 3. Iterate over nums from smallest to largest, keeping curr as current number
    for i1 := 0; i1 < len(nums2)-1; i1++ {
        curr := nums2[i1]

        // 4. Perform binary search for index of target-curr
        i2:= binarySearch(target-curr, nums2[i1+1:])

        if i2 > -1{
            // fmt.Println("i1",i1,"i2",i2,"curr",curr,"nums2[i2+i1+1]",nums2[i2+i1+1])
            c1 := containsIndex(nums2[i1],nums)
            c2 := containsIndex(nums2[i2+i1+1],nums)
            if c2 == c1 {
                c2 = containsIndex(nums2[i2+i1+1],nums[c1+1:]) + c1 + 1
            }
            // fmt.Println("c1;c2",c1,c2)
            return []int{c1, c2}
        }

    }

    return []int{0, 0}
}
