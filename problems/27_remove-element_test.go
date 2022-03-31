package problems

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	n := removeElement(nums, 3)
	if n != 2 {
		t.Errorf("error : %v, %d", nums, n)
	}
	nums = []int{3, 3}
	n = removeElement(nums, 3)
	if n != 0 {
		t.Errorf("error : %v, %d", nums, n)
	}
	nums = []int{2}
	n = removeElement(nums, 3)
	if n != 1 {
		t.Errorf("error : %v, %d", nums, n)
	}
}
