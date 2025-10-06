package stack

import "strconv"

// Time Complexity O(n)
// Space Complexity O(n)
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, token := range tokens {
		if op, isOperator := operators[token]; isOperator {
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			stack = append(stack, op(left, right))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
