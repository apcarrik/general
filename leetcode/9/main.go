func isPalindrome(x int) bool {

    // idea: find the number modulus 10 repeatedly. keep track of the digits you've seen so far
    
    // Disregard if negative
    if x < 0 {
        return false
    }

    // Create slice of individual digits
    var digits []int
    for x > 9 {
        digits = append(digits, x%10)
        x=x/10
    }
    digits = append(digits, x%10)

    // Check digits slice is palindrome
    for len(digits) > 1 {
        if digits[0] != digits[len(digits)-1] {
            return false
        }
        digits = digits[1:len(digits)-1]
    }

    return true

}
