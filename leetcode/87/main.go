func sameCharFreq(s1, s2 string) bool {
	// Check if s1 and s2 have the exact same character frequency (use frequency maps)
	s1map := map[rune]int{}
	s1mapKeyCount := 0
	for _,r := range s1 {
		if _,ok := s1map[r]; ok {
			s1map[r]++
		} else {
			s1map[r] = 1
			s1mapKeyCount++
		}
	}
	
	s2map := map[rune]int{}
	s2mapKeyCount := 0
	for _,r := range s2 {
		if _,ok := s2map[r]; ok {
			s2map[r]++
		} else {
			s2map[r] = 1
			s2mapKeyCount++
		}
	}

	if s1mapKeyCount != s2mapKeyCount {
		return false
	}

	for k,v := range s1map {
		if _,ok := s2map[k]; ok  {
			if s2map[k] != v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func isScrambleRec(s1, s2 string, lookupPtr *map[string]bool) bool {
	// check lookup for s1 and s2
	if _,ok := (*lookupPtr)[s1+"/"+s2]; ok {
		return (*lookupPtr)[s1+"/"+s2]
	}
	if _,ok := (*lookupPtr)[s2+"/"+s1]; ok {
		return (*lookupPtr)[s2+"/"+s1]
	}
	
	// check if s1 and s2 are the same
	same := true
	for i:=0; i<len(s1); i++ {
		if s1[i] != s2[i] {
			same  = false
		}
	}
	if same {
		(*lookupPtr)[s1+"/"+s2] = true
		return true
	}

	// try to find a splitting of s1 and s2 with the same characters in each split
	for i:=1; i<len(s1); i++{
		s1_1 := s1[:i]
		s1_2 := s1[i:]
		s2_1a := s2[:i]
		s2_2a := s2[i:]
		s2_1b := s2[len(s2)-i:]
		s2_2b := s2[:len(s2)-i]
		if sameCharFreq(s1_1, s2_1a) {
			if isScrambleRec(s1_1, s2_1a, lookupPtr) && isScrambleRec(s1_2, s2_2a, lookupPtr) {
				(*lookupPtr)[s1+"/"+s2] = true
				return true
			}

		}
		if sameCharFreq(s1_1, s2_1b) {
			if isScrambleRec(s1_1, s2_1b, lookupPtr) && isScrambleRec(s1_2, s2_2b, lookupPtr) {
				(*lookupPtr)[s1+"/"+s2] = true
				return true
			}
		}

	}
	// if no split possible, return false
	(*lookupPtr)[s1+"/"+s2] = false
	return false
    
}


func isScramble(s1 string, s2 string) bool {
	lookup := map[string]bool{}
	return  isScrambleRec(s1, s2, &lookup)
}
