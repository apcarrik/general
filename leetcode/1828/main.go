/* Notes:
*   - to check if point is within circle, I need to find distance between point and center of circle
*   - I think it would help to order points by x or y coordinate, or store them in a struct. Then I could
*   query only points from xmin to xmax based on circle radius and center.
*
*   Idea: Store points in map with x value as key and list of y coordinates as value. For each query, search
*   all x coordinates that could fall within the circle and search all y coordinates within those.
*/

import (
    "math"
)

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

    // Go through queries, for each searching any x coordinates in pointsMap that could be within the circle
    pointsInside := []int{}
    for j,query := range queries {
        xj := query[0]
        yj := query[1]
        rj := query[2]
        pointsInside = append(pointsInside,0)
        for x := xj - rj; x <= xj + rj; x++ {
            if _,ok := pointsMap[x]; ok {
                for _,y := range pointsMap[x] {
                    if distance([2]int{x,y}, [2]int{xj,yj}) <= float64(rj) {
                        pointsInside[j]++
                    }
                }
            }
        }
    }

    return pointsInside
}
