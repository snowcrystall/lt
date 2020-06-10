package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	n := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(n)
	n = lengthOfLongestSubstring("bbbbb")
	fmt.Println(n)
	n = lengthOfLongestSubstring("pwwkew")
	fmt.Println(n)
	n = lengthOfLongestSubstring("p")
	fmt.Println(n)
	n = lengthOfLongestSubstring(" ")
	fmt.Println(n)
}
