package test

import (
	"testing"
	"github.com/snowcrystall/leetcode-go/problems"
)
func TestUniqueMorseRepresentations(t *testing.T) {
	if 2 != problems.UniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}) {
		t.Errorf("error")
	}
}
