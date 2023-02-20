import "fmt"

func searchInsert(nums []int, target int) int {
    low := 0
    high := len(nums)
    mid := (high-low)/2
    for low < high {
        mid = (high-low)/2+low
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            low = mid+1
        } else {
            high = mid
        }
        fmt.Println("low:",low,"mid:",mid,"high:",high)
    }
    if nums[mid] < target {
        return mid + 1
    }
    return mid
    
}
