/* New Idea: First, modify all roads to have first city smaller than second. Then, sort all roads by first node, then 2nd in case of ties. Finally, iterativley create set of nodes connected to first starting from first node. Keep going until you find all connected nodes to first, and keep track of minimum distance in there.
    - also, keep a queue and add connections to nodes you find as you go along
        - use binary search to add new connection
*/

import (
    "sort"
)

type CitySortAsc [][]int

func (a CitySortAsc) Len() int {return len(a)}
func (a CitySortAsc) Swap(i,j int) {a[i], a[j] = a[j], a[i]}
func (a CitySortAsc) Less(i,j int) bool {
    if a[i][0] == a[j][0] {
        return a[i][1] < a[j][1]
    } else {
        return a[i][0] < a[j][0]
    }
}


type CitySortDesc [][]int

func (a CitySortDesc) Len() int {return len(a)}
func (a CitySortDesc) Swap(i,j int) {a[i], a[j] = a[j], a[i]}
func (a CitySortDesc) Less(i,j int) bool {
    if -a[i][0] == -a[j][0] {
        return -a[i][1] < -a[j][1]
    } else {
        return -a[i][0] < -a[j][0]
    }
}


// type QNode struct {
//     RoadIdx int
//     Next *QNode
// }

// func binSearch(target int, roads [][]int) int {
//     low := 0
//     high := len(roads)
//     for high > low {
//         mid := (high - low) / 2 + low
//         if roads[mid][0] >= target {
//             high = mid - 1
//         } else {
//             low = mid + 1
//         }
//     }
//     if low < len(roads){
//         if roads[low][0] < target && low < len(roads)-1 {
//             low++
//         }
//         if roads[low][0] == target {
//             return low
//         }
//     }

//     return -1 // if not found

// }

func minScore(n int, roads [][]int) int {
    minDist := 100001
    connectedCities := map [int]bool{1:true, n:true}
    connectedCitiesLen := 2
    oldConnectedCitiesLen :=0

    for true{

        // Modify roads to always have first city be smaller number
        for _,road := range roads {
            if road[0] > road[1] {
                road[0], road[1] = road[1], road[0]
            }
        }

        // Sort roads by first city then second city in case of ties
        sort.Sort(CitySortAsc(roads))

        // Search for all cities connected to city 1
        for _,road := range roads {
            city1, city2, dist := road[0], road[1], road[2]
            //fmt.Println("checking asc city1, city2:", city1, city2)
            if _,ok := connectedCities[city1]; ok {
                if  _,ok := connectedCities[city2]; !ok{
                    connectedCities[city2] = true
                    connectedCitiesLen++
                }
                if dist < minDist {
                    minDist = dist
                }
            }
            if _,ok := connectedCities[city2]; ok {
                if  _,ok := connectedCities[city1]; !ok{
                    connectedCities[city2] = true
                    connectedCitiesLen++
                }
                if dist < minDist {
                    minDist = dist
                }
            }
        }
        //fmt.Println("after asc, connectedCities:", connectedCities)
        // Now, do it all again but from city n and sort list in complete opposite direction    
        
        // Modify roads to always have first city be bigger number
        for _,road := range roads {
            if road[0] < road[1] {
                road[0], road[1] = road[1], road[0]
            }
        }

        // Sort roads by first city then second city in case of ties
        sort.Sort(CitySortDesc(roads))

        // Search for all cities connected to city 1
        for _,road := range roads {
            city1, city2, dist := road[0], road[1], road[2]
            //fmt.Println("checking desc city1, city2:", city1, city2)
            if _,ok := connectedCities[city1]; ok {
                if  _,ok := connectedCities[city2]; !ok{
                    connectedCities[city2] = true
                    connectedCitiesLen++
                }
                if dist < minDist {
                    minDist = dist
                }
            }
            if _,ok := connectedCities[city2]; ok {
                if  _,ok := connectedCities[city2]; !ok{
                    connectedCities[city1] = true
                    connectedCitiesLen++
                }
                if dist < minDist {
                    minDist = dist
                }
            }
        }
        ////fmt.Println("after desc, connectedCities:", connectedCities)
        if connectedCitiesLen == oldConnectedCitiesLen {
            break
        }
        oldConnectedCitiesLen = connectedCitiesLen

    }
    return minDist





    // // Create queue initialized to first road in roads, as its gaurenteed to contain city 1
    // roadQ := &QNode{
    //     RoadIdx: 0,
    // }
    // roadQTail := roadQ

    // // Search for all roads starting from city 1
    // for i:=1; i<len(roads); i++ {
    //     if roads[i][0] != 1 {
    //         break
    //     }
    //     newNode := &QNode{
    //         RoadIdx: i,
    //     }
    //     roadQTail.Next = newNode
    //     roadQTail = newNode
    // }

    // // Create map of non-searched nodes
    // citiesNotSearched := map[int]map[int]bool{}
    // for _,road := range roads {
    //     if _,ok := citiesNotSearched[road[0]]; !ok {
    //         citiesNotSearched[road[0]] = map[int]bool{road[1]:true}
    //     } else {
    //         citiesNotSearched[road[0]][road[1]] = true
    //     }
    // }

    // // Search through every city in roadQ, adding additional cities to queue as you go
    // minDistance := 100001
    // // roadIdxesSearched := map [int]bool{}
    // for roadQ != nil {
    //     road := roads[roadQ.RoadIdx]
    //     if _,ok := citiesNotSearched[road[0]]; ok {
    //         if _,ok := citiesNotSearched[road[0]][road[1]]; ok {
    //             city2 := road[1]
    //             // check if minDistance needs to be updated
    //             if road[2] < minDistance {
    //                 minDistance = road[2]
    //             }
    //             // search for index of first city2 road in roads and add all to roadQ
    //             firstIdx := binSearch(city2, roads)
    //             if firstIdx != -1 {
    //                 for firstIdx < len(roads) && roads[firstIdx][0] == city2 {
    //                     newNode := &QNode{
    //                         RoadIdx: firstIdx,
    //                     }
    //                     roadQTail.Next = newNode
    //                     roadQTail = newNode
    //                     firstIdx++
    //                 }
    //             }
    //         }
    //     }

    //     roadQ = roadQ.Next
    // }

    // return minDistance

}

/* New Idea: Create graph structure of cities, and then run DFS to find connections nodes to 1 and n. 
    - You can use recursion to implement DFS, but you'll need to pass the current set of connections cities with each call
    - You can also pass a minimum distance and continually update that

    Note: it still takes too long (!)
*/

// type Connection struct {
//     City *CityNode
//     Distance int
// }

// type CityNode struct {
//     ID int
//     Neighbors []*Connection
// }

// func findConnections(node *CityNode, minDist int, connectionsPtr *map[*CityNode]map[*CityNode]bool) int {
//     connections := *connectionsPtr
//     if _,ok := connections[node]; !ok {
//         connections[node] = map[*CityNode]bool{}
//     }
//     for _,conn := range node.Neighbors {
//         if _,ok := connections[node][conn.City]; !ok {
//             connections[node][conn.City] = true            
//             if conn.Distance < minDist {
//                 minDist = conn.Distance
//             }
//             // recursivley call findConnections on this city
//             minDist = findConnections(conn.City, minDist, &connections)
//         }
//     }


//     return minDist
// }

// func minScore(n int, roads [][]int) int {

//     cities := map [int]*CityNode{}
//     // 1. Create representation of graph
//     for _,road := range roads {
//         c1, c2, dist := road[0], road[1], road[2]
//         // add cities if not already present
//         if _,ok := cities[c1]; !ok {
//             cities[c1] = &CityNode{
//                 ID: c1,              
//             }
//         }
//         if _,ok := cities[c2]; !ok {
//             cities[c2] = &CityNode{
//                 ID: c2,                
//             }
//         }
//         // add road between cities
//         cities[c1].Neighbors = append(cities[c1].Neighbors, &Connection{City: cities[c2], Distance: dist})
//         cities[c2].Neighbors = append(cities[c2].Neighbors, &Connection{City: cities[c1], Distance: dist})
//     }

//     // 2. Do DFS to find all connections nodes to 1 and n
//     connections := map[*CityNode]map[*CityNode]bool{}
//     minDist := findConnections(cities[1], 10001, &connections)
//     minDist = findConnections(cities[n], minDist, &connections)
    
//     return minDist
// }



/* Notes:
    - Minimum score of path means minimum distance road in that path
    - Minimum possible score of patsh is the minimum distance road in any path between two cities
    - So, we're really looking for the minimum distance road that is part of any path between two cities.
    - Actually, we're looking for the minimum distance road that is connections to cities 1 and n
   Idea: Start by searching for all cities connections to 1 and n. Then, go through all roads and find the minimum that contains at least one city in the set connections to 1 and n
    Problem: it takes too long! In the worst case, you're searching over n roads in roads n times, so O(n^2). 
*/

// func minScore(n int, roads [][]int) int {

//     // 1. Keep a set of cities reachable by cities 1 and n. Move roads connecting these cities to another list
//     citiesReachable := map [int]bool{1:true, n:true}
//     oldRoadsLen := len(roads)
//     minDist := 10001

//     outerLoop:
//     for len(roads) > 0 {

//         for i:=0; i<len(roads); i++{
//             road := roads[i]
//             city1, city2, distance := road[0], road[1], road[2]
//             _,city1Reachable := citiesReachable[city1]
//             _,city2Reachable := citiesReachable[city2]

//             if city1Reachable || city2Reachable {

//                 // update minDist if applicable
//                 if distance < minDist {
//                     minDist = distance
//                 }

//                 // add new cities to citiesReachable
//                 if !city1Reachable {
//                     citiesReachable[city1] = true
//                 }
//                 if !city2Reachable {
//                     citiesReachable[city2] = true
//                 }

//                 // remove road from roads
//                 if i == len(roads) - 1 {                    
//                     roads = roads[:i]
//                 } else if i == 0 {
//                     roads = roads[1:]
//                 } else {
//                     roads = append(roads[:i], roads[i+1:]...)
//                 }

//                 i--

//             }

//         }

//         if len(roads) == oldRoadsLen {
//             break outerLoop
//         }

//         oldRoadsLen = len(roads)

//     }

//     return minDist
    
// }
