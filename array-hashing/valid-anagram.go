package arrayHashing

import (
	"sort"
)

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// Time Complexity O(n)
// Space Complexity O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true

	// Time Complexity O(n log n)
	// Space Complexity O(1)
	// return sortString(s) == sortString(t)
}
