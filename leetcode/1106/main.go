/* Notes:
    - Were basically creating a compiler
    - The usual way is to need to create a parser & tokenizer then apply the language rules
    - What we can probably do is use a state machine
    - Idea: find outermost expression, then recursivley call the inner subexpression and return the value of that with the operator applied
    - first character of expression will tell you what operator is
        - question: can first char be 't' or 'f' and have any subexpression? A:no
    - when parsing arguments, need to count parenthesis to not inaverdantly count recursive commas

*/

func parseArgs(expression string) []string {
    // input is a set of comma separated arguments, and output is array where each index is a separate argument as a string
    args := []string{}

    // // edge case: first char in expression is operator, so we know this will be one operator
    // switch expression[0] {
    //     case '!', '&', '|':
    //         return []string{expression}
    // }
    // isolate args separated by commas
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
        // if c == ',' {
        //     args = append(args, expression[argStart:i])
        //     argStart = -1
        // } else {
        //     if argStart == -1 {
        //         argStart = i
        //     }
        // }

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

    result := false

    // Find expression operator and operands
    args := []string{}
    switch expression[0] {
        case 't':
            result = true
        case 'f':
            result = false
        case '!':
            args = parseArgs(expression[2:len(expression)-1])
            result = !parseBoolExpr(args[0])            
        case '&':
            args = parseArgs(expression[2:len(expression)-1])
            result = true
            for _,arg := range args {
                result = result && parseBoolExpr(arg)
            }
        case '|':
            args = parseArgs(expression[2:len(expression)-1])
            result = false
            for _,arg := range args {
                result = result || parseBoolExpr(arg)
            }
    }


    // Return result
    return result
        
}
