func findKDigit(k int) (kDig, newK int) {
        kDig = 0
        if k>0 {
            kDig = k%10
            if k<10 {
                newK = 0
            } else {
                newK = k/10
            }
        }
    return
}

func addToArrayForm(num []int, k int) []int {
    retSlice := []int{}
    carry := 0
    kDig := 0
    for i:=0; i<len(num); i++{
        numDig := 0
        // Find k digit
        kDig, k = findKDigit(k)

        // Find num digit
        numDig = num[len(num)-1-i]

        // Append to retSlice
        newDig := kDig+numDig+carry
        carry = 0
        if newDig >= 10 {
            carry = 1
            newDig = newDig%10
        }
        retSlice = append([]int{newDig},retSlice...)

    }

    // iterate until k is not 0
    for k>0 {
        kDig, k = findKDigit(k)
        newDig := kDig+carry
        carry = 0
        if newDig > 9 {
            carry = 1
            newDig = newDig%10
        }
        retSlice = append([]int{newDig},retSlice...)
    }

    // check if carry is not 0
    if carry != 0 {        
        retSlice = append([]int{carry},retSlice...)        
    }
    return retSlice
    
}
