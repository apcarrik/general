type MatchNode struct {
    Index int
    Next *MatchNode
    Prev *MatchNode
}

func removeMatch(match, matchesHead *MatchNode) ( *MatchNode,  *MatchNode) {    
    // remove match from matches
    if match == matchesHead {
        matchesHead = matchesHead.Next
        if matchesHead != nil {
            matchesHead.Prev = nil
        }
    } else {
        match.Prev.Next = match.Next
        if match.Next != nil{
            match.Next.Prev = match.Prev
        }
    }
    return match, matchesHead
}

func strStr(haystack string, needle string) int {
    // Idea: loop through each rune of needle. For each needle rune, record indicies of all matches of substrings that include all needle runes upto and including this needle rune until you reach the end. Then, find the first index of matching substrings. Note that the number of matches will be weakly decreasing for each iteration of needle rune.
    hRunes := []rune{}
    for _,hRune := range haystack {
        hRunes = append(hRunes, hRune)
    }

    var matchesHead *MatchNode
    for nRuneIdx, nRune := range needle {
        if nRuneIdx == 0 {
            matchesTail := matchesHead
            for hRuneIdx, hRune := range hRunes {
                if hRune == nRune {
                    if matchesHead == nil {
                        matchesHead = &MatchNode{
                            Index: hRuneIdx,
                        }
                        matchesTail = matchesHead
                    } else {
                        matchesTail.Next = &MatchNode{
                            Index: hRuneIdx,
                            Prev: matchesTail,
                        }
                        matchesTail = matchesTail.Next
                    }
                }
            }
        } else {
            for match:=matchesHead; match != nil; match = match.Next {
                matchIdx := match.Index
                if matchIdx + nRuneIdx < len(hRunes){
                    hRune := hRunes[matchIdx + nRuneIdx]
                    if hRune != nRune {
                        match, matchesHead = removeMatch(match, matchesHead)
                    }
                } else {
                    match, matchesHead = removeMatch(match, matchesHead)
                }
            }
        }
    }
    if matchesHead == nil {
        return -1
    }
    return matchesHead.Index
}
