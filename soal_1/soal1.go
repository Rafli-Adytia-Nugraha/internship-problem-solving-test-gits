package main

import (
	"fmt"
	"strings"
)

func A000124Formatted(n int) string {
	if n <= 0 {
		return ""
	}

	parts := make([]string, 0, n)
	for k := 0; k < n; k++ {
		val := k*(k+1)/2 + 1
		parts = append(parts, fmt.Sprintf("%d", val))
	}

	return strings.Join(parts, "-")
}

func main() {
	inputs := []int{7, 1, 12}

	for _, n := range inputs {
		fmt.Printf("Input: %d\nOutput: %s\n\n", n, A000124Formatted(n))
	}
}
