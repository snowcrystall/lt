package problems

import "testing"

func TestSearchInsert(t *testing.T) {
	if searchInsert([]int{1, 3, 5, 6}, 5) != 2 {
		t.Errorf("expert 2")
	}
	if searchInsert([]int{1, 3, 5, 6}, 2) != 1 {
		t.Errorf("expert 1, %d", searchInsert([]int{1, 3, 5, 6}, 2))
	}
}
