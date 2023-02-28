func manhattanDistance(x1, y1, x2, y2 int) int {
    xDelta := x1-x2
    yDelta := y1-y2
    if xDelta < 0 {
        xDelta = 0-xDelta
    }
    if yDelta < 0 {
        yDelta = 0-yDelta
    }
    return xDelta + yDelta
}

func nearestValidPoint(x int, y int, points [][]int) int {
    closestPointIdx := -1
    closestPointManhattanDistance := int((1<<63)-1)
    for i,point := range points {
        var px int
        var py int
        for i,v := range point {
            if i==0 {
                px = v
            } else if i==1{
                py = v
            }
        }
        if px == x || py == y {
            manhattanDistance := manhattanDistance(x, y, px, py)
            if manhattanDistance < closestPointManhattanDistance {
                closestPointManhattanDistance = manhattanDistance
                closestPointIdx = i
            }
        }
    }
    return closestPointIdx
}
