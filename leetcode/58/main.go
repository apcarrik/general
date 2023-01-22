import "strings"
func lengthOfLastWord(s string) int {
    s2 := strings.TrimSpace(s)
    return len(s2) - strings.LastIndexByte(s2,' ') - 1
}
