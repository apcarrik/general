import (
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	// Sort satisfaction
	sort.Ints(satisfaction)

	// Create backwards rolling sum and like-time coefficients from satisfaction and keep track of maxLTC seen
	maxLTC := 0
	for i:= len(satisfaction) -2; i >= 0; i-- {
		if i == len(satisfaction)-2 {
			satisfaction[i] += 2*satisfaction[i+1]
		} else {
			satisfaction[i] += ((2*satisfaction[i+1]) - satisfaction[i+2])
		}
		if satisfaction[i] > maxLTC{
			maxLTC = satisfaction[i]
		}
	}

	return maxLTC

}
