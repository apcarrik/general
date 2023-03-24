/* 

*/
func checkOnesSegment(str string) bool {
    contigSegSeen := false
    prevCharOne := false
    for _,c := range str {
        if c == '1' {
            if contigSegSeen && !prevCharOne{
                return false
            }
            contigSegSeen = true
            prevCharOne = true
        } else {
            prevCharOne = false
        }
    }
	return true


}


// func checkOnesSegment(s string) bool {
    
// }
