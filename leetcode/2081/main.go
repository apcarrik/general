import (
    "math"
)

/* getNextBaseTenPalindrome notes
 1-9
 11, 22, 33...99
 101, 111, 121,...191
 202, 212, ... 292
 303, 313, ... 393
 404...494
 ...
 909...999
 1001, 1111, 1221, 1331...1991
 2002, 2112, 2222, 2332...2992
 Idea: if old is odd, increment middlemost digit (if even, middlemost 2 digits). If this/these are already 9, reset it/them to 0, and move out to surrounding pair of digits and increment them. If all digits of number are 9, then return 1xxx1 with x being the number of digits of old -1

*/

func getNextBaseTenPalindrome(old int64) int64 {
    // Idea: create array of digits of old number, and try to increment the middle-most digit(s)
    digits := []int8{}
    allDigits9 := true
    for tmp:=old; tmp>0; tmp/=int64(10) {
        digits = append(digits, int8(tmp % int64(10)))
        if tmp % int64(10) != int64(9) {
            allDigits9 = false
        }
    }

    if allDigits9 {
        return int64( math.Pow(float64(10), float64(len(digits))) ) + 1
    }

    oddCarry := false
    if len(digits) % 2 == 1 { // odd
        if digits[len(digits)/2] == 9 {
            oddCarry = true
            digits[len(digits)/2] = 0
        } else {
            digits[len(digits)/2]++
        }        
    }
    if oddCarry || len(digits) % 2 == 0 {
        for i := (len(digits)/2) - 1; i >= 0; i-- {
            if digits[i] == 9 {
                digits[i] = 0
                digits[len(digits) - 1 - i] = 0
            } else {
                digits[i]++
                digits[len(digits) - 1 - i]++
                break
            }
        }
    }

    // create new number
    new := uint64(0)
    for i,digit := range digits {
        new += uint64(digit) * uint64(math.Pow(float64(10), float64(i)))
    }
    if new > uint64(9223372036854775807) {
        fmt.Println("Error: new max > int64. old:",old,"new:", new)
    }
    return int64(new)
}

func isBaseKPalindrome(candidate int64, k int64) bool {
    
    digits := []int8{}
    for tmp:=candidate; tmp > 0; tmp /=int64(k) {
        digits = append(digits, int8(tmp%k))
    }
    for i:=0; i<len(digits)/2; i++{
        if digits[i] != digits[len(digits)-1-i] {
            return false
        }
    }
    return true
}

/* kMirror Notes
    - 0 does not count for any base
    - 1 is always the first k-mirror number
    - Idea: contruct the palindromes for base-k, and then for each check if it is palindrome for base-10
        - this may take too long, as it is an exhaustive search
        - actually, since k < 10, it would be faster to find all base-10 palindromes and check each against base-k
*/

func kMirror(k int, n int) int64 {
    // fmt.Printf("[")
    // old := int64(9009)
    // for i:=0 ; i<1000; i++{
    //     old = getNextBaseTenPalindrome(old)
    //     fmt.Printf("%d, ",old)
    // }
    // fmt.Printf("]")
    // return 0
    //fmt.Println("k:",k,"n:",n)
    sum := int64(1)
    found := 1
    candidate := int64(1)
    for found < n {
        candidate = getNextBaseTenPalindrome(candidate)
        if candidate == int64(-8446744082709551617) {
            //fmt.Println("candidate= -8446744082709551617", isBaseKPalindrome(candidate, int64(k)))
        }
        if isBaseKPalindrome(candidate, int64(k)){
            //fmt.Println(found,"Candidate found:", candidate)
            sum += candidate
            found++
        }
    }
    return sum    

}
