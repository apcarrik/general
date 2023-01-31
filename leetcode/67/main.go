
func addBinary(a string, b string) string {
    // Idea: starting from end of each string, add each digit one by one
    var maxlen int
    if len(a) > len(b) {
        maxlen = len(a)
    } else {
        maxlen = len(b)
    }
    aidx := len(a)-1
    bidx := len(b)-1
    retval := ""
    carry := 0
    for i := 0; i<maxlen; i+= 1  {
        adig := 0
        if aidx >= 0 {
            adig = int(a[aidx] - '0')
        }
        bdig := 0
        if bidx >= 0 {
            bdig = int(b[bidx] - '0')
        }
        if adig + bdig + carry == 3 {
            fmt.Println("sum=",adig + bdig + carry)
            retval = string('1') + retval
            carry = 1
        } else if adig + bdig + carry == 2 {
            retval = string('0') + retval
            carry = 1
        } else if adig + bdig + carry == 1 {
            retval = string('1') + retval
            carry = 0
        } else {
            retval = string('0') + retval
            carry = 0
        }     
        aidx -= 1
        bidx -= 1
    }
    if carry == 1 {
        retval = string('1') + retval
    }
    return retval
}
