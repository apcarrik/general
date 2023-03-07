/* Notes:
*   - to check if point is within circle, I need to find distance between point and center of circle
*   - I think it would help to order points by x or y coordinate, or store them in a struct. Then I could
*   query only points from xmin to xmax based on circle radius and center.
*
*   Idea: Store points in map with x value as key and list of y coordinates as value. For each query, search
*   all x coordinates that could fall within the circle and search all y coordinates within those.
*   
*   Optimization: Perform binary search on y values to find range suitable for match. Requires sorting y values
*/

import (
    "math"
    "sort"
)

func binarySearch(listPtr *[]int, x int) int {
    list := *listPtr
    low := 0
    high := len(list)
    for high > low {
        mid := (high-low)/2 + low
        switch {
            case list[mid] >= x:
                high = mid-1
            case list[mid] < x:
                low = mid+1
        }
    }
    return low

}

func distance(p1, p2 [2]int) float64 {
    deltaX := p1[0]-p2[0]
    if deltaX < 0 {
        deltaX = -deltaX
    }
    deltaY := p1[1]-p2[1]
    if deltaY < 0 {
        deltaY = -deltaY
    }
    return math.Sqrt(math.Pow(float64(deltaX),2) + math.Pow(float64(deltaY),2))
}

func countPoints(points [][]int, queries [][]int) []int {

    // Create map of x coordinates : list of y coordinates
    pointsMap := map [int][]int{}
    for _,point := range points {
        xi:=point[0]
        yi:=point[1]
        if _,ok := pointsMap[xi]; !ok {
            pointsMap[xi] = []int{}
        }
        pointsMap[xi] = append(pointsMap[xi], yi)
    }

    // Sort all y cordinates in pointsMap
    for x,_ := range pointsMap {
        sort.Ints(pointsMap[x])
    }

    // Create sorted list of x values
    xList := []int{}
    for k,_ := range pointsMap {
        xList = append(xList, k)
    }
    sort.Ints(xList)

    // Go through queries, searching for coordinates that could fit within (x,y) bounds of query circle
    pointsInside := []int{}
    for j,query := range queries {
        xj := query[0]
        yj := query[1]
        rj := query[2]
        pointsInside = append(pointsInside,0)
        // Binary search for x search range
        for xIdx := binarySearch(&xList,xj-rj);
            xIdx < len(xList) && xList[xIdx] <= xj+rj;
            xIdx++ {
            x := xList[xIdx]
            if ys,ok := pointsMap[x]; ok && len(ys) > 0 {
                // Binary search for y search range  
                for yIdx := binarySearch(&ys,yj-rj);
                    yIdx < len(ys) && ys[yIdx] <= yj+rj;
                    yIdx++ {
                    y := ys[yIdx]
                    if distance([2]int{x,y}, [2]int{xj,yj}) <= float64(rj) {
                        pointsInside[j]++
                    }
                }
            }
        }
    }

    return pointsInside
}
