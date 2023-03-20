/* Notes:
    - to plant a new flower, need at least 3 open spots in a row
    - for a section of open spots that's m long, you can plant up to (m-1)/2 flowers
    - Idea: iterate through entire flowerbet to find all sections of open spaces at least 3 long and sum how many flowers can be planted there, and compare with n.
    - Problem: beginning and end edge cases, for example flowerbed=[0,0,1,0,0]; n=2 would return false but should be true
        - idea: pad beginning and end with [1,0] and [0,1], respectivley

*/

func canPlaceFlowers(flowerbed []int, n int) bool {

    // pad beginning and end of flowerbed
    flowerbed = append(append([]int{1,0}, flowerbed...), []int{0,1}...)

    flowerSpots := 0
    emptyStart := -1
    for i,plot := range flowerbed {
        if plot == 0 {
            if emptyStart == -1 {
                emptyStart = i
            }
        } else {
            if emptyStart != -1 {
                flowerSpots += (i-emptyStart-1)/2
                emptyStart = -1
            }
        }
    }
    if n <= flowerSpots {
        return true
    }
    return false
}
