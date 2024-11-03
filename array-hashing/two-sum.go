package arrayHashing

// Time Complexity O(n)
// Space Complexity O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {

		if val, found := m[target-num]; found {
			return []int{val, i}
		}

		m[num] = i
	}
	return nil
}
