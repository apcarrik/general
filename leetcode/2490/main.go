func isCircularSentence(sentence string) bool {
    // Idea: loop through sentence, keeping track of last rune before space and matching it to first rune of next word

    var firstR, lastR rune
    newWord := false
    for i,r := range sentence {
        if i == 0 {
            firstR = r
            lastR = r
        } else {
            if r == rune(' ') {
                newWord = true
            } else {
                if newWord == true {
                    if r != lastR {
                        return false
                    }
                    newWord = false
                }
                lastR = r
            }
        }
    }
    if firstR != lastR{
        return false
    }
    return true
}
