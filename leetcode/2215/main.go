func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]bool{}
	for _,x := range nums1 {
		if _,ok := set1[x]; !ok {
			set1[x] = true
		}
	}
	set2 := map[int]bool{}
	for _,x := range nums2 {
		if _,ok := set2[x]; !ok {
			set2[x] = true
		}
	}

	difference := [][]int{[]int{}, []int{}}
	for k,_ := range set1 {
		if _,ok := set2[k]; !ok {
			difference[0] = append(difference[0], k)
		}
	}
	for k,_ := range set2 {
		if _,ok := set1[k]; !ok {
			difference[1] = append(difference[1], k)
		}
	}
	
	return difference
}
