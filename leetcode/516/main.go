type Bounds [2]int

func longestPalindromeSubseqRec(sPtr *string, low, high int , strMapPtr *map[string]int) int {
	strMap := *strMapPtr
	s := *sPtr
	longest := 1

	// check if in map already
	if _,ok := strMap[s[low:high+1]]; ok {
		return strMap[s[low:high+1]]
	}

	// check if low & high are match
	if low == high {
		strMap[s[low:high+1]] = 1
		return 1
	} else if high - low == 1 {
		if s[low] == s[high] {
			strMap[s[low:high+1]] = 2
			return 2
		} else {
			strMap[s[low:high+1]] = 1
			return 1
		}
	} else { // high - low >= 2

		// check for any match for low
		lowMatchLongest := 1
		for lowMatch := high; lowMatch > low; lowMatch-- {
			if s[low] == s[lowMatch] {
				if lowMatch - low == 1 {
					lowMatchLongest = 2
				} else {				
					newLow := low+1
					newHigh := lowMatch-1
					lowMatchLongest = longestPalindromeSubseqRec(&s, newLow, newHigh, strMapPtr) + 2
				}
				break
			}
		}

		// check for any match for high
		highMatchLongest := 1
		for highMatch := low; highMatch < high; highMatch++ {
			if s[high] == s[highMatch] {
				if high - highMatch == 1 {
					highMatchLongest = 2
				} else {				
					newLow := highMatch+1
					newHigh := high-1
					highMatchLongest = longestPalindromeSubseqRec(&s, newLow, newHigh, strMapPtr) + 2
				}
				break
			}
		}

		if lowMatchLongest > highMatchLongest {
			if lowMatchLongest > longest {
				longest = lowMatchLongest
			}
		} else {			
			if highMatchLongest > longest {
				longest = highMatchLongest
			}
		}


		tmpLongest := longestPalindromeSubseqRec(&s, low+1, high-1, strMapPtr)
		if tmpLongest > longest {
			longest = tmpLongest
		}


		strMap[s[low:high+1]] = longest


	}
	return longest
	
}

func longestPalindromeSubseq(s string) int {	
	strMapPtr := &map[string]int{}
	retval := longestPalindromeSubseqRec(&s, 0, len(s)-1, strMapPtr)
	return retval
}
