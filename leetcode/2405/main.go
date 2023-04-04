func partitionString(s string) int {
	chars := [26]int{}
	subs := 1
	for i:=0; i<len(s); i++ {
		ci := s[i]-97
		if chars[ci] > 0 {
			subs++
			for i := range chars {
				chars[i] = 0
			}
		}
		chars[ci]++		
	}
	return subs
}
