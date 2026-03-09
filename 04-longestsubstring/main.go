package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		if index, found := charIndex[s[right]]; found && index >= left {
			left = index + 1
		}

		charIndex[s[right]] = right

		windowLen := right - left + 1

		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcdabcdbb"))
}
