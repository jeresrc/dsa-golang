package arrayHashing

// Using Bucket Sort
// Time Complexity O(n)
// Space Complexity O(n)
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	countSliceLength := len(nums) + 1
	countSlice := make([][]int, countSliceLength)

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	var res []int

	for i := countSliceLength - 1; len(res) != k; i-- {
		res = append(res, countSlice[i]...)
	}

	return res
}
