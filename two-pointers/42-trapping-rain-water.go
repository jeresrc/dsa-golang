package twoPointers

func trap(heights []int) int {
	water := 0
	left, right := 0, len(heights)-1
	bigLeft, bigRight := 0, 0

	for left < right {

		if heights[left] < bigLeft {
			water += bigLeft - heights[left]
		} else {
			bigLeft = heights[left]
		}

		if heights[right] < bigRight {
			water += bigRight - heights[right]
		} else {
			bigRight = heights[right]
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return water
}
