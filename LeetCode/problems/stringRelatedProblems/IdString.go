package problems

import (
	"fmt"
	"strings"
)

func SmallestNumber(pattern string) string {

	stack := []int{}
	result := strings.Builder{}
	n := len(pattern)

	for i := 0; i <= n; i++ {
		stack = append(stack, i+1)

		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				result.WriteString(fmt.Sprintf("%d", stack[len(stack)-1]))
				stack = stack[:len(stack)-1]

			}
		}
	}
	return result.String()

}
