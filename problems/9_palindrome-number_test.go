package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Logf("%v isPalindrome %v", 1221, isPalindrome(1221))
	t.Logf("%v isPalindrome %v", -1221, isPalindrome(-1221))
	t.Logf("%v isPalindrome %v", 0, isPalindrome(0))
	t.Logf("%v isPalindrome %v", 12321, isPalindrome(12321))
}
