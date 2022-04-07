package test

import (
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	lcp := problems.LongestCommonPrefix(strs)
	if lcp != "fl" {
		t.Errorf("%v lcp is %s , expect fl", strs, lcp)
	}

	strs = []string{"dog", "racecar", "car"}
	lcp = problems.LongestCommonPrefix(strs)
	if lcp != "" {
		t.Errorf("%v lcp is %s , expect ''", strs, lcp)
	}

	strs = []string{"ab", "a", "abc"}
	lcp = problems.LongestCommonPrefix(strs)
	if lcp != "a" {
		t.Errorf("%v lcp is %s , expect 'a'", strs, lcp)
	}
}
