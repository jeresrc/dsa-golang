package twoPointer

import "sort"

func threeSum(nums []int) [][]int {
	// nums is an array of nums
	// res should contain triplets where sumatory == 0
	// res shouldn't contain duplicated triplets
	// => i can sort nums and check for i != i - 1
	// E res[i] => res[i] has 3 elements from nums, nums[j] + nums[k] + nums[l] = 0, j != k != l
	res := [][]int{}
	sort.Ints(nums)

	for i, candidate := range nums {

		if candidate > 0 {
			break
		}

		if i > 0 && candidate == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			threeSum := candidate + nums[left] + nums[right]
			if threeSum > 0 {
				right--
			} else if threeSum < 0 {
				left++
			} else {
				res = append(res, []int{candidate, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return res
}
