/* Notes
    - Normall storing words is efficiently done with a Trie, however the '.' wildcard search will complicate this
    - It seems like knowing the length of all words is important here
    - Since only 10k max words, it is probably better to store all words than use Trie (3 dots would open up to 17k searches)
    - We could store words in map with key: length, value: word as string
    - What if we created a 26 bit match array for each word, corresponding to if each letter is used?

ex. word: "applejack" -> bitmap[1,0,1...] -> a, c, e, j, k, l, p
    match: "apple" -> bitmap[1,0,0,...] -> a, e, l, p
*/

type WordDictionary struct {
    WordMap map [int][]string    
}


func Constructor() WordDictionary {
    return WordDictionary{map [int][]string{}}
}


func (this *WordDictionary) AddWord(word string)  {
    wordLen := len(word)
    if _,ok := this.WordMap[wordLen]; ok {
        for _,existingWord := range this.WordMap[wordLen] {
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
    } else {
        this.WordMap[wordLen] = []string{word}
    }
    return
}


func (this *WordDictionary) Search(word string) bool {
    wordLen := len(word)
    
    for _,existingWord := range this.WordMap[wordLen] {
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
