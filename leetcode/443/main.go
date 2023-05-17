func compress(chars []byte) int {
    // Idea: keep pointer of index of chars to replace and a pointer that iterates through characters until you hit a new character. 

    start, end := 0,1
    retStrLen := 0
    
    for end <= len(chars) {
        if end == len(chars) || chars[start] != chars[end] {
            // update character
            chars[retStrLen] = chars[start]

            // create character count and append to nums
            charCount := end-start
            charCountDigits := 0
            if charCount > 1 {
                charCountDigits = 1
                // get number of digits
                for i:=charCount; i>9; i++{
                    charCountDigits++
                    i/=10
                }

                // Place digits on nums
                for i:=charCountDigits; i>0; i-- {
                    chars[retStrLen+i] = string( charCount % 10 )[0] + 48
                    charCount /= 10
                }

            }
            retStrLen = retStrLen + 1 + charCountDigits
            start = end
        }
        end++
    }
    return retStrLen    
}
