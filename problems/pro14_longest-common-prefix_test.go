package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	lcp := longestCommonPrefix(strs)
	if lcp != "fl" {
		t.Errorf("%v lcp is %s , expect fl", strs, lcp)
	}

	strs = []string{"dog", "racecar", "car"}
	lcp = longestCommonPrefix(strs)
	if lcp != "" {
		t.Errorf("%v lcp is %s , expect ''", strs, lcp)
	}

	strs = []string{"ab", "a", "abc"}
	lcp = longestCommonPrefix(strs)
	if lcp != "a" {
		t.Errorf("%v lcp is %s , expect 'a'", strs, lcp)
	}
}
