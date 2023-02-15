import (
    "container/heap"
    "math"
)
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return -h[i] < -h[j] }
func (h IntHeap) Swap(i, j int) { h[i],h[j] = h[j],h[i] }
func (h *IntHeap) Pop() interface{} {
    retVal := (*h)[len(*h)-1]
    *h = (*h)[0:len(*h)-1]
    return retVal
}
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}



func maxKelements(nums []int, k int) int64 {
    // Attempted other much slower solutions but found the idea for a max heap in existing solutions.
    // Idea: order nums as max heap. For k iterations, on each iteration add max element to score. Then update that element with ceil(element/3) and place back on heap and call fix. Return sum after k iterations.

    // Create maxheap based on nums
    nHeap := &IntHeap{}
    *nHeap = append(*nHeap,nums...) // RUNTIME: amortized constant time for resizing, so O(n)
    heap.Init(nHeap) // RUNTIME: O(n)

    // Iterate k times, adding to sum and updaing elements
    sum := 0
    for k>0 {
        maxElem := heap.Pop(nHeap).(int) // RUNTIME: O(log(n))
        sum += maxElem
        heap.Push(nHeap, int(math.Ceil(float64(maxElem)/3)) ) // RUNTIME: O(log(n))
        k--
    }

    // return sum
    return int64(sum)
}

// Overall runtime: O(n + k log(n))
