//import "fmt"

// func getBeauty(sPtr *string, low, high int) int {
//     // Make frequency map of characters in string
//     s := (*sPtr)
//     freqMap := map [rune]int{}
//     for _,c := range s[low:high]{
//         if _,ok := freqMap[c]; ok {
//             freqMap[c]++
//         } else {
//             freqMap[c] = 1
//         }
//     }

//     // Find largest and smallest characters from map
//     max := 0
//     min := len(s)+1
//     for _,v := range freqMap {
//         if v > max {
//             max = v
//         }
//         if v < min {
//             min = v
//         }
//     }
//     return max-min
// }

// func beautySum(s string) int {
    
//     // Iterate through all substrings of s
//     beautySum := 0
//     for i:=len(s); i>1; i-- {
//         for j:=0; j<=len(s)-i; j++{
//             beautySum += getBeauty(&s, j, j+i)
//         }
//     } 
//     return beautySum
// }


// //-----------------
// func coalesceMaps(fMap1, fMap2 map [rune]int) *map [rune]int {
//     // iterate through every k,v in map2 and add to corresponding k,v in map 1
//     //fmt.Println("coalesce: fmap1, fmap2:", fMap1, fMap2)
//     newMap := map [rune]int{}
//     for k,v := range fMap1 {
//         newMap[k] = v
//     }

//     for k,v := range fMap2 {
//         if _,ok := fMap1[k]; ok {
//             newMap[k] += v
//         } else {
//             newMap[k] = v
//         }
//     }
//     //fmt.Println("new map:",newMap)
//     return &newMap

// }


// func getBeautyFast(fMap map [rune]int) int {    
//     // Find largest and smallest characters from map
//     max := 0
//     min := 501
//     for _,v := range fMap {
//         if v > max {
//             max = v
//         }
//         if v < min {
//             min = v
//         }
//     }
//     // fmt.Println("fMap:", fMap, "max,min:",max, min)
//     return max-min
// }

// func makeFreqMap(s string) *map [rune]int {
//     // Make frequency map of characters in string
//     freqMap := map [rune]int{}
//     for _,c := range s{
//         if _,ok := freqMap[c]; ok {
//             freqMap[c]++
//         } else {
//             freqMap[c] = 1
//         }
//     }
//     return &freqMap

// }

// func findIndicies(j,i int) ([2]int, [2]int) {
//     var freqMapIdx1, freqMapIdx2 [2]int
//     if i == 1 {
//         freqMapIdx1 = [2]int{j, j}
//         freqMapIdx2 = [2]int{j+i, j+i}
//     } else {
//         freqMapIdx1 = [2]int{j, j+(i/2)}
//         freqMapIdx2 = [2]int{j+(i/2)+1, j+i}
//     }
//     return freqMapIdx1, freqMapIdx2
// }

// func beautySum(s string) int {
    
//     // Iterate through all substrings of s
//     beautySum := 0
//     var fMapKey [2]int
//     freqMaps := map [[2]int]*map[rune]int{}
//     for i:=0; i<len(s); i++ {
//         if i>0 {
//             //fmt.Println(freqMaps)
//             // for idx,freqMap := range freqMaps {
//             //     fmt.Println(idx, freqMap)
//             // }
//         }
//         for j:=0; j<len(s)-i; j++{
//             if i== 0{
//                 fMapKey = [2]int{j, j+i}
//                 freqMaps[fMapKey] = makeFreqMap(s[j: j+i+1])
//                 //fmt.Println(fMapKey, "freqMaps[fMapKey]", freqMaps[fMapKey])
//             } else {
//                 freqMapIdx1,freqMapIdx2 := findIndicies(j,i)
//                 fMapKey = [2]int{j, j+i}
//                 //fmt.Println(freqMapIdx1,freqMapIdx2,fMapKey)
//                 freqMaps[fMapKey] = coalesceMaps(*freqMaps[freqMapIdx1],*freqMaps[freqMapIdx2])
//                 newBeauty := getBeautyFast(*freqMaps[fMapKey])
//                 //fmt.Println(freqMapIdx1,freqMapIdx2,fMapKey,freqMaps[fMapKey], newBeauty)
//                 beautySum += newBeauty
//             }
//         }
//     } 
//     return beautySum
// }

// -------------------

// With help from https://leetcode.com/problems/sum-of-beauty-of-all-substrings/solutions/2590973/python-simplified-solution-94-faster/
func getBeauty(fMap map [rune]int) int {
    max := 0
    min := 501
    for _,freq := range fMap {
        if freq > max {
            max = freq
        }
        if freq < min {
            min = freq
        } 
    }
    return max-min
}

func beautySum(s string) int {
    sRunes := []rune{}
    for _,rune := range s {
        sRunes = append(sRunes, rune)
    }
    
    // Iterate through all substrings of s
    beautySum := 0
    for i:=0; i<len(sRunes)-1; i++ {
        fMap := map [rune]int{}
        for j:=i; j<len(sRunes); j++{
            if _,ok := fMap[sRunes[j]]; ok {
                fMap[sRunes[j]]+=1
            } else {
                fMap[sRunes[j]] = 1
            }
            beautySum += getBeauty(fMap)
        }
    } 
    return beautySum
}
