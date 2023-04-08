func isIdealPermutation(nums []int) bool {
    // Basically, the problem is asking of all gobal inversions are local inversions. I.E. is every decrease as you go right only to the next index
    // Idea: if any value is more than 2 greater than it's index, or less than 2 less than its index, then return false

    for i:=0; i<len(nums); i++{
        if nums[i] > i+1 || nums[i] < i-1{
            return false
        }
    }
    return true
}
