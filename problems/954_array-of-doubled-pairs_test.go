package problems

import "testing"

func TestCanReorderDoubled(t *testing.T) {
	if !canReorderDoubled([]int{3, 6, 3, 6}) {
		t.Errorf("error [3,6,3,6]")
	}
	if !canReorderDoubled([]int{4, -2, 2, -4}) {
		t.Errorf("error [4,-2,2,-4]")
	}
	if canReorderDoubled([]int{2, 1, 2, 6}) {
		t.Errorf("error [2,1,2,6]")
	}
	if canReorderDoubled([]int{-5, -2}) {
		t.Errorf("error [-5,-2]")
	}
}
