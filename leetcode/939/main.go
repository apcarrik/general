func Abs(x int) int {
    if x<0{
        return -x
    }
    return x
}

func minAreaRect(points [][]int) int {
    // Thoughts:
    //  Need to check if any 4 points form a rectangle. Maybe sort all points by x value, then by y value?
    //  Then search for all rectangles, keeping track of the minimum area seen? Actually, i'll need two hashmaps,
    //  one for x_val as key and one for y_val as key.
    // Idea: 1. Make hashmap with key=x_val and value= array of y values, and another with y_val as key and array of x_val after that. Then iterate through each x value, exhaustivley searching for all possible rectangles, and keeping track of the minimum area seen.

    // Edge case: points has <4 points
    if len(points) < 4 {
        return 0
    }
    maximumArea := 40000*40000+1
    minimumArea := maximumArea // initialized to maximum_area + 1

    // Create map of x values with a list of their corresponding y values
    xMap := map[int][]int{}
    for _,point := range points{
        x := point[0]
        y := point[1]
        _,exists := xMap[x]
        if exists{
            xMap[x] = append(xMap[x],y)
        } else {
            xMap[x] = []int{y}
        }
    }

    // Create map of y values with a map of their corresponding x values
    yMap := map[int]map[int]bool{}
    for _,point := range points{
        x := point[0]
        y := point[1]
        _,exists := yMap[y]
        if exists{
            yMap[y][x] = true
        } else {
            yMap[y] = map[int]bool{x:true}
        }
    }

    // Go through each x value, cross-checking each corresponding y values for possible rectangles
    for x1,xMapYs := range xMap {
        if len(xMapYs) > 1 { // can't make a rectangle if there is only 1 point at that x coordinate
            // check every possible combination of 2 y values
            for y1i,y1 := range xMapYs {
                for y2i := y1i + 1; y2i < len(xMapYs); y2i++ {
                    // check if exists rectangle with [x1, y1], [x1, y2], [x2, y1], [x2, y2]
                    y2 := xMapYs[y2i]
                    if len(yMap[y1]) > 1 && len(yMap[y2]) > 1 {
                        for x2, _ := range yMap[y1] {
                            if x2 != x1 {
                                _, x2ExistsForY2 := yMap[y2][x2]
                                if x2ExistsForY2 {
                                    // rectangle found, update minimum area if it is smaller
                                    newRectArea := Abs(x1-x2)*Abs(y1-y2)
                                    if newRectArea < minimumArea {
                                        minimumArea = newRectArea
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    if minimumArea < maximumArea{
        return minimumArea
    } else {
        return 0
    }

}
