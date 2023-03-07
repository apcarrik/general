/* Notes
    - bus with shortest trip time will be limiting factor
    - max trip time would be totalTrips*shortestBusTripTime
    - Idea: could I use binary search to search between 1 and totalTrips for the shortest bus trip time increments and see if total trips of all busses in that time interval are > totalTrips
        - calculating trips for every bus will be O(n) for n busses, and binary search takes O(n) time
    - It is too slow. Is there a better way to speed this up?
        - Idea: find average trip time and divide total time by that

*/

import (
    "math"
    //"fmt"
)
func getTripsCompleted(tripTimesPtr *[]int, totalTime int) int {
    tripTimes := *tripTimesPtr
    tripsCompleted := 0
    // calculate how many total trips can be completed in totalTime
    for _, tripTime := range tripTimes {
        tripsCompleted += totalTime / tripTime
    }

    return tripsCompleted
}

func getAvgTripsPerMinTripTime(tripTimesPtr *[]int, minTripTime float64) float64 {
    tripTimes := *tripTimesPtr
    avgTrips := float64(0)
    // calculate how many total trips can be completed in totalTime
    for _, tripTime := range tripTimes {
        avgTrips += minTripTime / float64(tripTime)
    }
    return avgTrips
}


func minimumTime(tripTimes []int, totalTrips int) int64 {
    
    // find minimum trip time of all times
    minTripTime := 10000001
    for _,tripTime := range tripTimes {
        if tripTime < minTripTime {
            minTripTime = tripTime
        }
    }

    avgTripPerMinTripTime := getAvgTripsPerMinTripTime(&tripTimes, float64(minTripTime))
    //fmt.Println(getTripsCompleted(&tripTimes, 3))
    //fmt.Println(getTripsCompleted(&tripTimes, 4))
    //fmt.Println(getTripsCompleted(&tripTimes, 5))
    // perform binary search to find minimum total time to achieve totalTrips over all busses
    minTime := int( math.Floor(float64(totalTrips) / avgTripPerMinTripTime)*0.9 )*minTripTime
    maxTime := int( math.Ceil(float64(totalTrips) / avgTripPerMinTripTime)*1.1  )*minTripTime
    //fmt.Println(minTripTime, avgTripPerMinTripTime, minTime, maxTime, getTripsCompleted(&tripTimes, minTime),getTripsCompleted(&tripTimes, maxTime))
    for maxTime > minTime {
        medianTime := (maxTime - minTime)/2 + minTime
        numTripsCompleted := getTripsCompleted(&tripTimes, medianTime)
        //fmt.Println(medianTime, numTripsCompleted)
        if numTripsCompleted >= totalTrips {
            maxTime = medianTime
        } else {
            minTime = medianTime+1
        }

    }

    // minTime now contains the result of binary search. It may be wrong, however, since minTime and maxTime are
    // approximations, so the following searches are needed
    if getTripsCompleted(&tripTimes, minTime) >= totalTrips{
        for getTripsCompleted(&tripTimes, minTime) >= totalTrips {
            minTime--
        }
        return int64(minTime+1)
    } else if getTripsCompleted(&tripTimes, minTime) < totalTrips {
        for getTripsCompleted(&tripTimes, minTime) < totalTrips {
            minTime++
        }
    }
    return int64(minTime)

}
