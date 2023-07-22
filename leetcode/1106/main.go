/* Notes:
    - Were basically creating a compiler
    - The usual way is to need to create a parser & tokenizer then apply the language rules
    - What we can probably do is use a state machine
    - Idea: find outermost expression, then recursivley call the inner subexpression and return the value of that with the operator applied
    - first character of expression will tell you what operator is
        - question: can first char be 't' or 'f' and have any subexpression? A:no
    - when parsing arguments, need to count parenthesis to not inaverdantly count recursive commas
    - optimization: after recieving args list from parseArgs, do a quick check to see if it is a single 't'/'f' value and return that value instead of recursivley calling parseBoolExpr again.

*/

func parseArgs(expression string) []string {
    // input is a set of comma separated arguments, and output is array where each index is a separate argument as a string
    args := []string{}
    argStart := -1
    openParens := 0
    for i,c := range expression {
        switch c {
        case '(':
            openParens++
        case ')':
            openParens--
        case ',':
            if openParens == 0 {
                args = append(args, expression[argStart:i])
                argStart = -1
            }
        default:
            if argStart == -1 {
                argStart = i
            }
        }

    }
    // edge cases
    if argStart == -1 && len(args) == 0 {
        // no commas found, so entire expression is argument
        args = append(args, expression)
    } else if argStart != -1 {
        // one last arg
        args = append(args, expression[argStart:len(expression)])
    }
    return args
}


func parseBoolExpr(expression string) bool {

    // Find expression operator and operands
    switch expression[0] {
        case 't':
            return true
        case 'f':
            return false
        case '!':
            if len(expression) == 4 { // optimization
                return expression[2] == 'f' 
            } else {
                return !parseBoolExpr(expression[2:len(expression)-1])
            }                        
        case '&':
            args := parseArgs(expression[2:len(expression)-1])
            if len(args) == 1 && len(args[0]) == 1 { // optimization
                return args[0][0] == 't' 
            } else {
                result := true
                for _,arg := range args {
                    result = result && parseBoolExpr(arg)
                }
                return result
            }
        case '|':
            args := parseArgs(expression[2:len(expression)-1])
            if len(args) == 1 && len(args[0]) == 1 { // optimization
                return args[0][0] == 't' 
            } else {
                result := false
                for _,arg := range args {
                    result = result || parseBoolExpr(arg)
                }
                return result
            }
    }

    return false
        
}
