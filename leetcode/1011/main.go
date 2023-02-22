import (
    "math"
    //"fmt"
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

    // Continually increase minShipCapacity by 1 until all weights will ship
    for minShipCapacity < totalWeight {
        daysToShip := 1
        shipWeight := 0
        //fmt.Println("minShipCapacity=",minShipCapacity)
        for _,weight := range weights {
            shipWeight += weight
            if shipWeight > minShipCapacity {
                //fmt.Println(">Ship can hold through index:",i)
                daysToShip++
                shipWeight = weight
            }
            if daysToShip > days {
                break
            }
        }
        //fmt.Println("Days to ship:", daysToShip)
        if daysToShip <= days{
            break
        }
        minShipCapacity += 1
    }
    return minShipCapacity

}
