func partitionString(s string) int {
	chars := [26]bool{}
	subs := 1
	for i:=0; i<len(s); i++ {
		ci := s[i]-97
		if chars[ci] {
			subs++
			for i := range chars {
				chars[i] = false
			}
		}
		chars[ci] = true		
	}
	return subs
}
