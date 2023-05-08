import (
    "container/heap"
)
type MaxHeap []int

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {return -h[i] < -h[j]}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[0: len(*h)-1]
    return x
}

func minimumDeviation(nums []int) int {
    // Thoughts: 
    // - We can divide even elements by 2 until they are an odd number. Even elements let us reduce their value.
    // - Odd elements can only be multiplied once, and then they turn even.
    // - You can't reduce the largest odd element and you can't increase the smallest even element. So, when your largest element is odd and smallest is even, you are at the minimum deviation and you can return.

    // Original Idea: Greedy solution. Copy all numbers in 2 heaps: one min and one max. Find the deviation between the min and max and store this as deviation. Try to manipulate the max to reduce it and then the min to increase it. Update the deviation, keeping track in the change of deviations. If the deviations stop changing between iterations, return the deviation as minimum deviation.

    // Improved Solution: From the solutions section, singhabhinash gave a great concise write-up. The core element I was missing was that we can turn all odd numbers into even by multiplying them by 2, and then we only have to reduce even numbers, requiring only one max heap. We can just keep the minimum of the heap in a single variable.

    // Remove duplicates from nums
    numsMap := map [int]bool{nums[0]: true}
    for _, num := range nums {
        _,exists := numsMap[num]
        if !exists{
            numsMap[num] = true
        }
    }

    // Store unique values from nums into numsUnique
    numsUnique := []int{}
    for num,_ := range numsMap {
        numsUnique = append(numsUnique,num)
    }

    // Initialize and populate maxHeap with only even numbers, and find minimum
    min := 1000000001
    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)
    for _, num := range numsUnique {
        if num % 2 == 1 { // multiply odd numbers by 2 to make them even
            num*=2
        }
        heap.Push(maxHeap, num)
        if num < min {
            min = num
        }
    }

    // Keep track of deviation
    minDeviation := (*maxHeap)[0] - min

    for true {
        max := heap.Pop(maxHeap).(int)
        if max-min < minDeviation {
            minDeviation = max-min
        }
        if max %2 == 1 {
            break
        }
        max /= 2
        if max < min {
            min = max
        }
        heap.Push(maxHeap, max)
    }

    return minDeviation

}
