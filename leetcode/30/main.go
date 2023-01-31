import (
    "strings" 
    //"fmt"
    )
func findSubstring(s string, words []string) []int {
        // Idea: create a moving window that searches for permutation concatenation on every n length substring, where n is the word length * number of words
        // question: can words contain duplicate words? Assume no. (actually, yes)
        wordlen := len(words[0])
        numwords := len(words)
        n := numwords * wordlen
        var retvec []int
        for i := 0; i<=len(s)-n; i+=1{
            // Check each word-length subsubstring in substring, and check if it matches any words
            var wordsFound []bool
            for k:=0; k<numwords; k++{
                wordsFound = append(wordsFound,false)
            }
            //fmt.Println("substr:",s[i:i+n])
            for j:=i; j < i+n; j+=wordlen { // check each possible word position
                //fmt.Println(" subsubstr:", s[j:j+wordlen])
                for w:=0; w<numwords; w+=1 { // check each word if it equals substring at that position
                    if wordsFound[w] == false && strings.Contains(s[j:j+wordlen], words[w])  {
                        //fmt.Println(" > word:",words[w]," found")
                        wordsFound[w] = true
                        break
                    } 
                }
            }
            isConcatenatedPermutation := true
            for w:=0; w<len(wordsFound); w+=1{
                if wordsFound[w] == false{
                    isConcatenatedPermutation = false
                }
            }            
            if isConcatenatedPermutation {
                retvec = append(retvec, i)
            }
        }
        return retvec
    
}
