type Trie struct {
    Letter [52]*Trie
    End [52]bool
}


func Constructor() Trie {
    return Trie{}
}

func charBytes(char byte) int8 {    
    // Turn character bytes into index for Trie
    charIndex := char
    if charIndex >= 97 && charIndex <=122 { // lowercase letter
        charIndex -= 97
    } else { // prefix bytes is not an english alphabet character
        return int8(-1)
    }
    return int8(charIndex)
}


func (this *Trie) Insert(word string)  {
    // Idea: add new character if not already on Trie, and then recursivley call with that letter's trie and remaining letters of word
    
    if len(word) == 0 || this == nil{
        return
    }    

    // Turn character bytes into index for Trie
    charIndex := int(charBytes(word[0]))
    if charIndex == -1 { // check if character is not recognized english character
        return
    }    
    
    if len(word) == 1 {
        // set end if this is last character in word
        this.End[charIndex] = true 
    } else {
        // otherwise, recursivley call Insert on this letter's Trie with remaining letters in word
        if this.Letter[charIndex] == nil {
            this.Letter[charIndex] = &Trie{}
        }
        this.Letter[charIndex].Insert(word[1:])
    }

    return
    
}


func (this *Trie) Search(word string) bool {
    // Idea: start with first character of prefix, and if found recursivley search for next characters using this characters Trie
    if len(word) == 0 {
        return false
    }

    // Turn character bytes into index for Trie
    charIndex := int(charBytes(word[0]))
    if charIndex == -1 {
        return false
    }

    if len(word) == 1 {
        // return End if this is last character in word
        return this.End[charIndex]
    } else {
        // otherwise, recursivley call Search on this letter's Trie with remaining letters in word
        if this.Letter[charIndex] != nil {
            return this.Letter[charIndex].Search(word[1:])
        }
    }
    return false
}


func (this *Trie) StartsWith(prefix string) bool {
    // Idea: start with first character of prefix, and if found recursivley search for next characters using this characters Trie
    if len(prefix) == 0 {
        return false
    }

    // Turn character bytes into index for Trie
    charIndex := int(charBytes(prefix[0]))
    if charIndex == -1 {
        return false
    }

    // Check if letter is in trie, either as prefix or end of word
    if this == nil || ( this.Letter[charIndex] == nil && this.End[charIndex] == false ){
        return false
    } else {
        if len(prefix) == 1 && (this.End[charIndex] == true || this.Letter[charIndex] != nil ){            
            return true
        } else if this.Letter[charIndex] != nil{            
            return this.Letter[charIndex].StartsWith(prefix[1:])
        }
    }
    return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

 /* Notes:
    - tries are usually trees with descendants for every letter (26 total in english). You store words by traversing the trie from the first letter to the last of the word.
    - data struct will probably have just array of pointers to next trie node. I'm going to use values 0-25 to represent letters a-z, and 26-51 to represent letters A-Z
    - Insert basically uses search, so i'll implement search first
    - Also, we'll need to denote if any letter at any level of trie is end of word. So make another 52 element array to denote if letter corresponding to that index is ending a word.
    - Question: can a prefix be an entire word?

*/
