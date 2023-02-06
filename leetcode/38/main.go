import (
    "strings"
    "strconv"
)
func countAndSay(n int) string {
    // Idea: start from 1 and iterate until you hit n
    saying := "1"
    for i:=1; i<n; i++{
        // Iterate through each char in saying
        var prevChar byte
        prevCharCount := 0
        var newSaying string
        for j:=0; j<len(saying); j++ {
            if saying[j] == prevChar {
                prevCharCount += 1
            } else {
                if prevCharCount > 0 {
                    newSaying += strconv.Itoa(prevCharCount) + string(rune(prevChar))
                }
                prevChar = saying[j]
                prevCharCount = 1
            }
        }        
        newSaying += strconv.Itoa(prevCharCount) + string(rune(prevChar))
        saying = newSaying
    }
    return saying
}
