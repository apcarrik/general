func validateStackSequences(pushed []int, popped []int) bool {
	pushPtr := 0
	popPtr := 0
	stack := []int{}
	popLoop:
	for popPtr <= len(popped)-1 {
		pop := popped[popPtr]
		if len(stack) > 0 && stack[len(stack)-1] == pop {
            // just pop this pop element from stack
            stack = stack[:len(stack)-1]
            popPtr++
        } else {
            // push to stack until you hit pop element, or run out of elements to push
            for pushPtr <= len(pushed)-1 {
                push := pushed[pushPtr]
                pushPtr++
                if push == pop {
                    popPtr++
                    continue popLoop
                } else {
                    stack = append(stack, push)
                }
            }
            return false
        }
	}	
    return true
}
