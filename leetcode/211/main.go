/* Notes
    - Normall storing words is efficiently done with a Trie, however the '.' wildcard search will complicate this
    - It seems like knowing the length of all words is important here
    - Since only 10k max words, it is probably better to store all words than use Trie (3 dots would open up to 17k searches)
    - We could store words in map with key: length, value: word as string

    Optimizations:
    - What if we created a 26 bit mask array for each word, corresponding to if each letter is used?
    - Note, we would need it to be a partial match, as match could have wildcard '.' characters

ex. word:  "apple" -> bitmap[1,0,0,0,1,0,0,0,0,0,0,1,0,0,0,1,...] -> a, e, l, p
    match: "app.e" -> bitmap[1,0,0,0,1,0,0,0,0,0,0,0,0,0,0,1,...] -> a, e, p
    Operator: XOR  -> bitmap[0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,...] -> word XOR match
    Operator: AND  -> bitmap[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,...] -> match AND (word XOR match)
*/

type WordDictionary struct {
    WordMap map [int][]string    
    BitMask map [int][]int32 // 26 bit int
}


func Constructor() WordDictionary {
    return WordDictionary{
        WordMap: map [int][]string{},
        BitMask: map [int][]int32{},
    }
}

func createBitmask(word string) int32 {
    bitmask := int32(0)
    for i:=0; i<len(word); i++{
        char := word[i]
        if char != '.'{
            bitmask = bitmask | 1<<(char-97)
        }
    }
    return bitmask
}

func checkBitmask(word, match int32) bool {
    if match & (word ^ match) == 0 {
        return true
    }
    return false
}


func (this *WordDictionary) AddWord(word string)  {
    // return
    wordLen := len(word)
    newBitmask := createBitmask(word)
    if _,ok := this.WordMap[wordLen]; ok {
        for j,existingWord := range this.WordMap[wordLen] {
            // Optimization: check character bitmask first
            if !checkBitmask(this.BitMask[wordLen][j], newBitmask) {
                continue
            }
            
            sameWord := true
            for i:=0; i<len(word); i++{
                if !(word[i] == '.') && word[i] != existingWord[i]{
                    sameWord = false
                    break
                }
            }
            if sameWord {
                return
            }
        }
        this.WordMap[wordLen] = append(this.WordMap[wordLen], word)
        this.BitMask[wordLen] = append(this.BitMask[wordLen], newBitmask)
    } else {
        this.WordMap[wordLen] = []string{word}
        this.BitMask[wordLen] = []int32{newBitmask}
    }
    return
}


func (this *WordDictionary) Search(word string) bool {
    wordLen := len(word)
    
    matchBitmask := createBitmask(word)
    for j,existingWord := range this.WordMap[wordLen] {
        // Optimization: check character bitmask first
        if !checkBitmask(this.BitMask[wordLen][j], matchBitmask) {
            continue
        }

        sameWord := true
        for i:=0; i<len(word); i++{
            if !(word[i] == '.') && word[i] != existingWord[i]{
                sameWord = false
                break
            }
        }
        if sameWord {
            return true
        }
    }
    return false
    
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
