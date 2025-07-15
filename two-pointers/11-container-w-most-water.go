package twoPointers

import "github.com/jeresc/dsa-golang/utils"

func MaxArea(heights []int) int {
	// heights is an array of bars
	// heights[i] is the height of the ith bar

	// if heights[j] and heights[k] exist => |k-j| = distance between bars
	// area between k'th bar and j'th bar = min(heights[j], heights[k]) * |k-j|

	// should return max area possibly obtainable between two bars

	leftPointer, rightPointer := 0, len(heights)-1
	res := 0

	for leftPointer < rightPointer {

		distance := rightPointer - leftPointer
		area := utils.Min(heights[leftPointer], heights[rightPointer]) * distance

		if area > res {
			res = area
		}

		if heights[leftPointer] < heights[rightPointer] {
			leftPointer += 1
		} else {
			rightPointer -= 1
		}

	}
	return res
}
