func reverse(x int) int {
    // Idea: find the number of digits in x, then create reverse x one digit at a time, checking for overflow at last digit    
    reversed := int32(0)
    lastDigit := int8(0)
    for X := int32(x); X != 0; X /= 10 {
        lastDigit = int8(X % 10)
        reversed = reversed * 10 + int32(lastDigit)
    }
    if int8(reversed % 10) != lastDigit {
        return 0
    }
    return int(reversed)
}
