package problems

import "testing"

func TestNumArray(t *testing.T) {
	nums := []int{7, 2, 7, 2, 0}
	numArr := ConstructorNumArray(nums)
	numArr.Update(4, 6)
	numArr.Update(0, 2)
	numArr.Update(0, 9)
	if 6 != numArr.SumRange(4, 4) {

		t.Errorf("error SumRange(4, 4)")
	}
	numArr.Update(3, 8)
	if 32 != numArr.SumRange(0, 4) {
		t.Errorf("error SumRange(0, 4)")
	}
	numArr.Update(4, 1)
	if 26 != numArr.SumRange(0, 3) {
		t.Errorf("error SumRange(0, 3)")
	}
	if 27 != numArr.SumRange(0, 4) {
		t.Errorf("error SumRange(0, 4)")
	}
}
