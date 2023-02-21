func singleNonDuplicate(nums []int) int {
    // Thoughts: If every integer in array appeared twice, we would know that every odd-indexed element would have a new integer to its right. So, wherever the single integer is, it will reverse that trend and the odd-indexed element will have a new integer to it's right. So, we need to find out where that occurs. Also, we know that array will always be pf odd length.
    // Idea: perform Binary Search since it's a sorted array.If both left and right adjacent to current integer are different, return that integer. Otherwise, if the above pattern is reversed at the searching location, search to the left. Else, search to the right. 

    // edge case, n=1
    if len(nums)==1{
        return nums[0]
    }
    lower := 0
    upper := len(nums)-1
    mid := 0
    for upper - lower > 2{
        mid = (upper-lower)/2+lower

        // check if this is the single integer
        if nums[mid-1] != nums[mid] && nums[mid] != nums[mid+1] {
            return nums[mid]
        }

        // otherwise, figure out which side to focus on
        if mid%2==1 { // mid is odd
            if nums[mid-1] == nums[mid] { // singleton is to the right
                lower = mid+1
            } else{ // singleton is to the left
                upper = mid-1
            }        
        } else { // mid is even
            if nums[mid-1] == nums[mid] { // singleton is to the left
                upper = mid
            } else{ // singleton is to the right
                lower = mid
            }        
        }
    }
    // only mid and above/below it are left; decide what to return
    mid = lower+1
    if nums[mid-1] == nums[mid] {
        return nums[mid+1]
    } else if nums[mid] == nums[mid+1] {
        return nums[mid-1]
    }
    return nums[mid]
}
