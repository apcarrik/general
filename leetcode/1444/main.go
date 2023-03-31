// func ways(pizza []string, k int) int {
    
// }

func numApples(rowMin, rowMax, colMin, colMax int, rowMapPtr *map[int][]int) int {
    rowMap := *rowMapPtr

    // return number of apples in pizza bounded by row/col min/max
    apples := 0
    for ri := rowMin; ri <= rowMax; ri++{
        if _,ok := rowMap[ri]; ok {
            for _,ci := range rowMap[ri]{
                if ci >= colMin && ci <= colMax {
                    apples++
                }
            }
        }
    }
    return apples
    
}

func hashBounds(k, rowMin, rowMax, colMin, colMax int) int {
    // return hash of coordinates
    return k + rowMin*100 + rowMax*10000 + colMin*1000000 + colMax*100000000
}

func waysRecursive(k, rowMin, rowMax, colMin, colMax int, rowMapPtr, colMapPtr *map[int][]int, memoPtr *map[int]int) int {
    // return number of ways to slice pizza maintaining >= 1 apple in slice and >= k-1 apples in smaller pizza
    memo := *memoPtr
    // Check if this k already covered in memo table
    boundsHash := hashBounds(k, rowMin, rowMax, colMin, colMax)
    if _,ok := memo[boundsHash]; ok {
        //fmt.Println(k, rowMin, rowMax, colMin, colMax, "returning:", memo[boundsHash])
        return memo[boundsHash]
    }

    // Base case: k=0, so if any apple in pizza return 1, else return false
    if k == 0 {
        if numApples(rowMin, rowMax, colMin, colMax, rowMapPtr) >= 1 {
            //fmt.Println(k, rowMin, rowMax, colMin, colMax, "returning: 1")
            return 1
        } else {
            //fmt.Println(k, rowMin, rowMax, colMin, colMax, "returning: 0")
            return 0
        }
    }

    possibleSlices := 0
    // check all possible row slices
    for ri:= rowMin+1; ri<=rowMax; ri++{
        // check if slice at row i would meet conditions, and recurse if so
        if numApples(rowMin, ri-1, colMin, colMax, rowMapPtr) > 0 && numApples( ri, rowMax, colMin, colMax, rowMapPtr) >= k-1 {
            //fmt.Println("checking:", k-1, ri, rowMax, colMin, colMax)
            possibleSlices += waysRecursive(k-1, ri, rowMax, colMin, colMax, rowMapPtr, colMapPtr, memoPtr)
        }
    }

    // check all possible col slices
    for ci:= colMin+1; ci<=colMax; ci++{
        // check if slice at row i would meet conditions, and recurse if so
        if numApples(rowMin, rowMax, colMin, ci-1, rowMapPtr) > 0 && numApples( rowMin, rowMax, ci, colMax, rowMapPtr) >= k-1 {
            //fmt.Println("checking:", k-1, rowMin, rowMax, ci, colMax)
            possibleSlices += waysRecursive(k-1, rowMin, rowMax, ci, colMax, rowMapPtr, colMapPtr, memoPtr)
        }
    }

    // update memo table
    memo[boundsHash] = possibleSlices

    // return sum of possible row and col slices
    //fmt.Println(k, rowMin, rowMax, colMin, colMax, "main returning:",possibleSlices)

    return possibleSlices
    

}


func ways( pizza []string, k int) int {
    // Create memo table
    memo := map[int]int{}

    // Create rowMap & colMap
    rowMap := map[int][]int{}
    colMap := map[int][]int{}
    for ri, row := range pizza {
        for ci, cell := range row {
            if cell == 'A' {
                if _,ok := rowMap[ri]; ok {
                    rowMap[ri] = append(rowMap[ri], ci)
                    colMap[ci] = append(colMap[ci], ri)
                } else {
                    rowMap[ri] = []int{ci}
                    colMap[ci] = []int{ri}
                }
            }
        }
    }

    // call recursive function with initial conditions and return result
    return waysRecursive(k-1, 0, len(pizza), 0, len(pizza[0]), &rowMap, &colMap, &memo) % 1000000007	

}
