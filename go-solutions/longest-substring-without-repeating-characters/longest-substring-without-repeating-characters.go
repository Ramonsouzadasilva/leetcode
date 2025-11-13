package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	start, maxLen := 0, 0

	for end := 0; end < len(s); end++ {
		char := s[end]

		if idx, found := lastSeen[char]; found && idx >= start {
			start = idx + 1
		}

		lastSeen[char] = end
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) 
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   
	fmt.Println(lengthOfLongestSubstring(""))           
}
