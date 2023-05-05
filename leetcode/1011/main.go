import (
    "math"
)

func shipWithinDays(weights []int, days int) int {
    // Idea: sum all weights. Divide by days. This gives average load per day. Find the maximum weight of all packages. The greater of maximum weight and average load per day will be minimum ship size. But, it may need to be larger. So, iterate through all weights to see if the minimum ship size will work. If not, increase size by 1 and try again.

    // find average load per day and maxWeight
    totalWeight := 0
    maxWeight := 0
    for _,weight := range weights{
        totalWeight += weight
        if weight > maxWeight {
            maxWeight = weight
        }
    }
    avgLoadPerDay := int(math.Ceil(float64(totalWeight)/float64(days)))

    // Determine starting minimum ship capacity
    minShipCapacity := avgLoadPerDay
    if maxWeight > avgLoadPerDay {
        minShipCapacity = maxWeight
    }

    // Optimization: use binary search instead, between minShipCapacity and maxWeight to find lowest minShipCapacity that will ship all items
    low := minShipCapacity
    mid := low
    high := totalWeight
    canShip := true
    for low < high {
        mid = (high-low)/2 + low

        // check if mid can ship
        daysToShip := 1
        shipWeight := 0
        for _,weight := range weights {
            shipWeight += weight
            if shipWeight > mid {
                daysToShip++
                shipWeight = weight
            }
            if daysToShip > days {
                break
            }
        }

        // change low or high bounds and canShip flag
        if daysToShip <= days{
            canShip = true
            high = mid
        } else {
            canShip = false
            low = mid +1
        }
    }

    // return correct minShipCapacity based on canShip flag
    if canShip {
        minShipCapacity = mid
    } else {
        minShipCapacity = mid + 1
    }
    return minShipCapacity

}
