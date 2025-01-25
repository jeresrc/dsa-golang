package stack

// Time Complexity O(4^n/sqrt(n))
// Space Complexity O(n)
func GenerateParenthesis(n int) []string {
	res := make([][]string, n+1)
	res[0] = []string{""}

	for k := 1; k <= n; k++ {
		res[k] = make([]string, 0)
		for i := 0; i < k; i++ {
			for _, left := range res[i] {
				for _, right := range res[k-i-1] {
					res[k] = append(res[k], "("+left+")"+right)
				}
			}
		}
	}

	return res[n]
}
