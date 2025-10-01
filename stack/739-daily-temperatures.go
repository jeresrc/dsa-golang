package stack

// Time Complexity O(n)
// Space Complexity O(n)
func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{} // store indices

	for idx, t := range temperatures {
		for len(stack) > 0 {
			lastIdx := stack[len(stack)-1]
			lastTemp := temperatures[lastIdx]

			if t <= lastTemp {
				break
			}

			stack = stack[:len(stack)-1]

			res[lastIdx] = idx - lastIdx
		}
		stack = append(stack, idx)
	}

	return res
}
