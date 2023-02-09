import "fmt"

func jump(nums []int) int {
    // Using official solution
    // Idea: use greedy approach to split problem up into ranges reachable at each step.

    n := len(nums)

    curFar := 0
    curEnd := 0
    answer := 0
    for i:=0; i<n-1; i++ {
        if curFar < i + nums[i]{
            curFar = i+nums[i]
        }
        if i == curEnd {
            answer++
            curEnd = curFar
        }
    }
    return answer

}
