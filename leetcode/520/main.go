func detectCapitalUse(word string) bool {

    allCharsCaps := true
    noCharsCaps := true
    numCaps := 0

    for _,c := range word {
        if c >= 0x41 && c<= 0x5a { // uppercase
            noCharsCaps = false
            numCaps ++
        } else if c >= 0x61 && c <= 0x7a { // lowercase
            allCharsCaps = false
        }
    }

    firstLetterCaps := false
    if word[0] >= 0x41 && word[0] <= 0x5a {
        firstLetterCaps = true
    }

    if allCharsCaps || noCharsCaps || (firstLetterCaps && numCaps == 1) {
        return true
    }
    return false
    
}
