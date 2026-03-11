package main

import (
	"fmt"
)

func findNonRepeatingCharacter(s string) int {
	count := map[rune]int{}

	for _, c := range s {
		count[c]++
	}

	for i, c := range s {
		if count[c] == 1 {
			return i
		}
	}

	return -1
}

func findNonRepeatingCharacter1(s string) rune {
	freq := map[rune]int{}

	for _, c := range s {
		freq[c]++
	}

	for _, c := range s {
		if freq[c] == 1 {
			return c
		}
	}
	return 0
}

func main() {
	fmt.Println(findNonRepeatingCharacter("Google"))          // G → index 0
	fmt.Println(findNonRepeatingCharacter1("Google"))         // G → index 0
	fmt.Println(string(findNonRepeatingCharacter1("Google"))) // G → index 0
}
