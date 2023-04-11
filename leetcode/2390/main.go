func removeStars(s string) string {
	starCount := 0
	newStr := ""
	for i:=len(s)-1; i >= 0; i--{
		switch s[i] {
        case '*':
            starCount++
        default:
            if starCount > 0 {
                starCount--
            } else {
                newStr = string(s[i]) + newStr
            }
		}
    }
    return newStr
}
