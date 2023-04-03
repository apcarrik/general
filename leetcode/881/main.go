import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	// sort people
	sort.Ints(people)

	// iterate through people, adding largest to boat and checking if smallest can come too. Keep pointers to largest and smallest
	largest := len(people)-1
	smallest := 0
	boats := 0
	for largest >= smallest {
		if people[smallest] <= limit - people[largest]  {
			smallest++
		}
		boats++
		largest--
	}

	return boats
}
