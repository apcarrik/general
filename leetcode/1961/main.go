func isPrefixString(s string, words []string) bool {
    sIdx := 0
    wordsIdx := 0
    for sIdx <= len(s)-1 {
        if wordsIdx > len(words)-1 {
            return false
        }
        wordIdx := 0
        for wordIdx < len(words[wordsIdx]) {
            if sIdx >= len(s) {
                return false
            }
            if s[sIdx] != words[wordsIdx][wordIdx] {
                return false
            }
            wordIdx++
            sIdx++
        }
        wordsIdx++
    }
    return true
}
