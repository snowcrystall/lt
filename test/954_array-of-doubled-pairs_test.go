package test

import (
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestCanReorderDoubled(t *testing.T) {
	if !problems.CanReorderDoubled([]int{3, 6, 3, 6}) {
		t.Errorf("error [3,6,3,6]")
	}
	if !problems.CanReorderDoubled([]int{4, -2, 2, -4}) {
		t.Errorf("error [4,-2,2,-4]")
	}
	if problems.CanReorderDoubled([]int{2, 1, 2, 6}) {
		t.Errorf("error [2,1,2,6]")
	}
	if problems.CanReorderDoubled([]int{-5, -2}) {
		t.Errorf("error [-5,-2]")
	}
}
