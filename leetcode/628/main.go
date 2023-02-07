import "sort"

func maximumProduct(nums []int) int {
    sort.Ints(nums)
    if len(nums) == 3 {
        return nums[0]*nums[1]*nums[2]
    }
    if len(nums) == 4 {
        // find which are positive
        num_positives := 0
        for i:=0; i<4; i+=1 {
            if nums[i] >= 0 {
                num_positives += 1
            }
        }
        switch num_positives {
            case 4:
                // find largest 3
                return nums[1]*nums[2]*nums[3]
            case 3:
                // return product of positive 3
                return nums[1]*nums[2]*nums[3]
            case 2:
                // return product of two negative times largest positive
                return nums[0]*nums[1]*nums[3]
            case 1:
                // return product of smallest two negative times the one positive
                return nums[0]*nums[1]*nums[3]
            default: 
                // all negative, return product of largest 3
                return nums[1]*nums[2]*nums[3]
                
        }
    }
    // Idea: Find largest 3 numbers and smallest 2 numbers, then check if 3 largest multiplied together is greater than smallest 2 multiplied with largest (in the case the two smallest are negative but have larger abs value than other 3 largest).
    if nums[0]*nums[1]*nums[len(nums)-1] > nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1] {
        return nums[0]*nums[1]*nums[len(nums)-1]
    } else {
        return nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1]
    }

}
