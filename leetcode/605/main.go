/* Notes:
    - to plant a new flower, need at least 3 open spots in a row
    - for a section of open spots that's m long, you can plant up to (m-1)/2 flowers
    - Idea: iterate through entire flowerbet to find all sections of open spaces at least 3 long and sum how many flowers can be planted there, and compare with n.
    - Problem: beginning and end edge cases, for example flowerbed=[0,0,1,0,0]; n=2 would return false but should be true
        - idea: pad beginning and end with [1,0] and [0,1], respectivley

*/

func canPlaceFlowers(flowerbed []int, n int) bool {

    flowerSpots := 0
    emptyStart := 1
    for i:=2; i<len(flowerbed)+4; i++ {
        var plot int
        switch i {
            case len(flowerbed)+2:
                plot = 0
            case len(flowerbed)+3:
                plot = 1
            default:
                plot = flowerbed[i-2]
        }
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
