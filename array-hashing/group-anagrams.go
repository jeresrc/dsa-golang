package arrayHashing

// Time Complexity O(m*n)
// Being m the length of the longest string and n the length of the array
// Space Complexity O(m)
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		anagramMap[count] = append(anagramMap[count], s)
	}

	result := make([][]string, len(anagramMap))
	i := 0

	for _, v := range anagramMap {
		result[i] = v
		i++
	}
	return result
}
