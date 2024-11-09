package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	start, maxlen := 0, 0
	for i, char := range s {
		if idx, ok := charMap[char]; ok && idx >= start {
			start = idx + 1
		}
		charMap[char] = i
		maxlen = max(maxlen, i-start+1)
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abac"))
}
