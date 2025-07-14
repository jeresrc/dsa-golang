package arrayHashing

// Hash Set
// Time Complexity O(n)
// Space Complexity O(n)
func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longest := 0
	for num := range numSet {
		if _, found := numSet[num-1]; !found {
			length := 1
			for {
				if _, exists := numSet[num+length]; exists {
					length++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

// Hash Map
// Time Complexity O(n)
// Space Complexity O(n)
func LongestConsecutive(nums []int) int {
	mp := make(map[int]int)
	res := 0

	for _, num := range nums {
		if mp[num] == 0 {
			left := mp[num-1]
			right := mp[num+1]

			sum := left + right + 1

			mp[num] = sum
			mp[num-left] = sum
			mp[num+right] = sum

			if sum > res {
				res = sum
			}
		}
	}

	return res
}
