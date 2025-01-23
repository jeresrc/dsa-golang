package arrayHashing

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
