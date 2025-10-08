package main

import (
	"fmt"
)

func replaceAt(s string, i int, r rune) string {
	return s[:i] + string(r) + s[i+1:]
}

func makePalindrome(s string, left, right, k int) (string, int, bool) {
	if left >= right {
		return s, k, true
	}

	if s[left] == s[right] {
		return makePalindrome(s, left+1, right-1, k)
	}

	if k <= 0 {
		return s, k, false
	}

	var bigger byte
	if s[left] > s[right] {
		bigger = s[left]
	} else {
		bigger = s[right]
	}

	s = replaceAt(s, left, rune(bigger))
	s = replaceAt(s, right, rune(bigger))
	return makePalindrome(s, left+1, right-1, k-1)
}

func maximizePalindrome(s string, left, right, k int) string {
	if left > right || k <= 0 {
		return s
	}

	if s[left] == '9' && s[right] == '9' {
		return maximizePalindrome(s, left+1, right-1, k)
	}

	if left == right && k > 0 {
		s = replaceAt(s, left, '9')
		return s
	}

	if s[left] != '9' || s[right] != '9' {
		need := 0
		if s[left] != '9' {
			need++
		}
		if s[right] != '9' {
			need++
		}
		if k >= need {
			s = replaceAt(s, left, '9')
			s = replaceAt(s, right, '9')
			return maximizePalindrome(s, left+1, right-1, k-need)
		}
	}
	return maximizePalindrome(s, left+1, right-1, k)
}

func highestPalindrome(s string, k int) string {
	pal, rem, ok := makePalindrome(s, 0, len(s)-1, k)
	if !ok {
		return "-1"
	}

	result := maximizePalindrome(pal, 0, len(pal)-1, rem)
	return result
}

func main() {
	input1, k1 := "3943", 1
	fmt.Printf("Input: %s\nk: %d\nOutput: %s\n\n", input1, k1, highestPalindrome(input1, k1))

	input2, k2 := "932239", 2
	fmt.Printf("Input: %s\nk: %d\nOutput: %s\n\n", input2, k2, highestPalindrome(input2, k2))

	input3, k3 := "12321", 1
	fmt.Printf("Input: %s\nk: %d\nOutput: %s\n", input3, k3, highestPalindrome(input3, k3))
}
