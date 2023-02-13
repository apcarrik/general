func countAsterisks(s string) int {
    // Idea: iterate through string, keeping track of asterisk count and if vertical bar phrases are open or closed
    asterisks := 0
    insideBarPhrase := false
    for i:=0; i<len(s); i++{
        switch s[i] {
            case '*':
                if !insideBarPhrase {
                    asterisks += 1
                }
            case '|':
                insideBarPhrase = !insideBarPhrase
        }
    }
    return asterisks    
}
