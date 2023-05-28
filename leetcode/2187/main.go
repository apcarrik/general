/*  Idea: Find a good approximation for the answer to get us in the "ballpark." Then, use binary search to search
    between max/min ballpark values and check if total trips of all busses in that time interval are > totalTrips.

    Note: Repeated attempts found that a min factor of 0.9 and max factor of 1.1 worked for most test cases, but 
    there were some low-valued test cases this didn't work for (i.e. case [9,3,10,5] and totalTime=2), so the
    last section is a last-ditch effort to search in increments of 1 if binary search didn't work.
*/

import (
    "math"
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

    // perform binary search to find minimum total time to achieve totalTrips over all busses
    // note the 0.9 and 1.1 factors because min/max times are approximations
    avgTripPerMinTripTime := getAvgTripsPerMinTripTime(&tripTimes, float64(minTripTime))
    minTime := int( math.Floor(float64(totalTrips) / avgTripPerMinTripTime)*0.9 )*minTripTime
    maxTime := int( math.Ceil(float64(totalTrips) / avgTripPerMinTripTime)*1.1  )*minTripTime
    for maxTime > minTime {
        medianTime := (maxTime - minTime)/2 + minTime
        numTripsCompleted := getTripsCompleted(&tripTimes, medianTime)
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
