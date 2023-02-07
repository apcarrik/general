import "strconv"

func maximumValue(strs []string) int {
    // Idea: Iterate through each str. For each, try to convert from string to int, and if that fails, get the length
    maxval := 0
    for i:=0; i<len(strs); i++ {
        val, err := strconv.Atoi(strs[i])
        if err != nil {
            val = len(strs[i])
        }
        if val > maxval {
            maxval = val
        }
    }
    return maxval
    
}
