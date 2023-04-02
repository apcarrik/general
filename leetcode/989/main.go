import (
    "math/big"
)

func addToArrayForm(num []int, k int) []int {
    // convert num to int and store as bigInt in retVal
    idx := int64(0)
    retVal := big.NewInt(0)
    pow10 := big.NewInt(0)
    for i:=len(num)-1; i>=0; i--  {
        pow10.Exp(big.NewInt(10),big.NewInt(idx),big.NewInt(0))
        retVal.Add(big.NewInt(0).Mul(big.NewInt(int64(num[i])),pow10),retVal)
        idx++
    }
    // add k to retVal
    retVal.Add(big.NewInt(int64(k)), retVal)

    // convert retVal to slice
    retSlice := []int{}
    newDigit := big.NewInt(0)
    for retVal.Cmp(big.NewInt(0)) == 1 { // todo: update this
        retVal.DivMod(retVal, big.NewInt(10), newDigit) // todo: update this
        retSlice = append([]int{int(newDigit.Int64())},retSlice...)
    }

    return retSlice    
}
