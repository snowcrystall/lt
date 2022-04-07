package test

import (
	"fmt"
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	n := problems.LengthOfLongestSubstring("abcabcbb")
	fmt.Println(n)
	n = problems.LengthOfLongestSubstring("bbbbb")
	fmt.Println(n)
	n = problems.LengthOfLongestSubstring("pwwkew")
	fmt.Println(n)
	n = problems.LengthOfLongestSubstring("p")
	fmt.Println(n)
	n = problems.LengthOfLongestSubstring(" ")
	fmt.Println(n)
}
