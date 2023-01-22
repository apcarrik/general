func lengthOfLastWord(s string) int {
    // Idea: starting from end, go character by character searching for last word
    // - Step 1: search for non-space character
    // - Step 2: start counter
    // - Step 3: end counter when you encounter space character or beginning of string
    wordlen := 0
    for i := len(s)-1; i >= 0; i-- {
        if wordlen == 0 {
            if s[i] != ' ' {
                wordlen += 1
            }
        } else {
            if s[i] == ' ' {
                return wordlen
            } else {
                wordlen += 1
            }
        }
    }
    return wordlen
}
