func isValid(s string) bool {
    // Idea: keep stack going from left to right of open left parens, once you hit right parens it must complement whatever is on top of stack else return false
    var stk []byte
    for i:=0; i<len(s); i++{
        switch s[i] {
        case ')' :
            if len(stk) < 1 || stk[len(stk)-1] != '(' {
                return false
            } else {
                stk = stk[:len(stk)-1]
            }
        
        case '}' :
            if len(stk) < 1 || stk[len(stk)-1] != '{' {
                return false
            } else {
                stk = stk[:len(stk)-1]
            }
        
        case ']' :
            if len(stk) < 1 || stk[len(stk)-1] != '[' {
                return false
            } else {
                stk = stk[:len(stk)-1]
            }
        
        case '(',  '{',  '[' :
            stk = append(stk, s[i])
        default:
            return false
        }
        
    }
    if len(stk) > 0 {
        return false
    } else {
        return true
    }
}
