/* Notes: 
    - we can break down the problem into single swaps, like bubble sort
    - since we only care about digits, we can treat bytes as runes
    - Idea: iterate through t string, starting from end index. Try to get the correct number there from s through bubble swaps.
    Then, go to next smallest index, etc until you finish the string

*/

import (
    "strconv"
)

type DLLNode struct {
    val int
    prev *DLLNode
    next *DLLNode
}

// type DLL struct {
//     head *DLLNode
//     tail *DLLNode
// }

func isTransformable(s string, t string) bool {

    // for digit in t, from end to beginning, try to use swaps to place the correct digit from s there.
    // note: there may be duplicate digits, so use the closest digit (i.e. the closest digit to the current index trying to be placed that hasn't been placed yet)
    // problem: runtime complexity is bad. In worst case, you'll be swapping every element the entire length of the array, 
    // so it is O(n^2)
    // Idea: instead of storing strings as rune array, store as doubly-linked list. This makes updating / shifting easier
    
    // Populate linked lists
    n := len(s)
    sdig,_ := strconv.Atoi(string(s[0]))
    sh := &DLLNode{ // s head
        val: sdig,
    }
    st := sh // s tail
    tdig,_ := strconv.Atoi(string(t[0]))
    th := &DLLNode{ // t head
        val: tdig,
    }
    tt := th // t tail
    for i:=1; i<n; i++ {
        sdig,_ := strconv.Atoi(string(s[i]))
        st.next = &DLLNode{
            val: sdig,
            prev: st,
        }
        st = st.next
        tdig,_ := strconv.Atoi(string(t[i]))
        tt.next = &DLLNode{
            val: tdig,
            prev: tt,
        }
        tt = tt.next
    }
    //fmt.Printf("s:[")
    for stmp:=sh; stmp != nil; stmp = stmp.next{
        //fmt.Printf("%d,",stmp.val)
    }
    //fmt.Printf("]\n")
    //fmt.Printf("t:[")
    for ttmp:=th; ttmp != nil; ttmp = ttmp.next{
        //fmt.Printf("%d,",ttmp.val)
    }
    //fmt.Printf("]\n")

    // Iterate through t list
    ti := tt // t iterator
    se := st // s end
    for ti != th {
        // find element to swap in s
        matchFound := false
        //fmt.Println(se)
        rangeMax := se.val
        for si := se; si != nil; si = si.prev {
            //fmt.Println("si.val:",si.val, "ti.val:",ti.val)
            if si.val == ti.val {                
                if rangeMax > si.val {
                    // can't bubble
                    //fmt.Println("can't move ", si.val, "ahead of ", rangeMax)
                    return false
                }
                // swap si and se if they are different nodes
                if si != se {
                    if si.next == se { // swap
                        if si == sh {
                            sh = se
                        } else {
                            si.prev.next = se
                        }          
                        if se == st {
                            st = si
                        } else {
                            se.next.prev = si
                        }
                        se.prev = si.prev      
                        si.next = se.next                        
                        se.next = si
                        si.prev = se
                        //fmt.Println("swapped", si, "and", se, "| sh=", sh, ", st=", st)
                    } else { // bubble si up above se
                        if si == sh {
                            sh = si.next
                            si.next.prev = nil
                        } else {
                            si.prev.next , si.next.prev = si.next, si.prev                            
                        }
                        if se == st {
                            st = si
                            si.next = nil
                        } else {
                            se.next.prev = si
                            si.next = se.next
                        }
                        se.next, si.prev = si, se
                        //fmt.Println("bubbled", si, "above", se, "| sh=", sh, ", st=", st)
                    }
                    se = si
                    //fmt.Println("now se = si:", se, si)
                    
                    //fmt.Printf("s:[")
                    for stmp:=sh; stmp != nil; stmp = stmp.next{
                        //fmt.Printf("%d,",stmp.val)
                    }
                    //fmt.Printf("]\n")
                }
                matchFound = true
                break
            } else {
                if rangeMax < si.val {
                    rangeMax = si.val
                }
            }
        }
        if !matchFound && se != nil{
            //fmt.Println("no match found")
            return false
        }

        // move on to prev element
        ti = ti.prev
        se = se.prev
    }

    // now, all swaps have been made. check that sh == th
    if sh.val != th.val {
        //fmt.Println("sh != th")
        return false
    }
    //fmt.Println("returning true\n------\n")
    return true

    // below was too slow, on account of bubbling through the entire array

    // ////fmt.Println("baS:",baS)
    // baT := []rune{}
    // for _,b := range t {
    //     baT = append(baT, b)
    // }
    // ////fmt.Println("baT:",baT)
    // for ti := len(baT) - 1; ti >= 0; ti--{
    //     dt := baT[ti]
    //     rangeMax,_ := strconv.Atoi(string(baS[ti]))
    //     matchFound := false
    //     for si := ti; si >= 0; si--{
    //         ds := baS[si]
    //         intS,_ := strconv.Atoi(string(ds))
            
    //         if ds == dt {
    //             if rangeMax > intS {
    //                 // can't bubble
    //                 ////fmt.Println(rangeMax, intS, baS)
    //                 return false
    //             }
    //             // bubble ds back to dt
    //             ////fmt.Println("before bubbling, baS:", baS)
    //             for si2:= si + 1; si2 <= ti; si2++ {
    //                 baS[si], baS[si2] = baS[si2], baS[si]
    //                 si = si2
    //             }                
    //             ////fmt.Println("after bubbling, baS:", baS)
    //             matchFound = true
    //             break
    //         } else {
    //             if rangeMax < intS {
    //                 rangeMax = intS
    //             }
    //         }
    //     }
    //     if !matchFound {
    //         return false
    //     }
    // }
    // return true
    
}

/*
s ="34521", t = "23415"
34215
32415
23415
*/

/*
s = "84532", t = "34852"
84352
43852
34852

*/
