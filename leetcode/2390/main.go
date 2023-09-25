func removeStars(s string) string {
	starCount := 0
	lastIdxWritten := len(s)
	for i:=len(s)-1; i >= 0; i--{
		switch s[i] {
			case '*':
				starCount++
			default:
				if starCount > 0 {
					starCount--
					if starCount == 0 {
						if lastIdxWritten == len(s) {
							s = s[:i]
						} else {
							s = s[:i] + s[lastIdxWritten:]
						}
                        lastIdxWritten = i
					}
				} else {
					lastIdxWritten = i
                }
		}
    }
    return s

}
