package problems

import "testing"

func TestPlusOne(t *testing.T) {
	res := plusOne([]int{1, 2, 3})
	if !testEq(res, []int{1, 2, 4}) {
		t.Errorf("[1,2,3] plusOne: %v", res)
	}
	res = plusOne([]int{1, 2, 9})
	if !testEq(res, []int{1, 3, 0}) {
		t.Errorf("[1,2,3] plusOne: %v", res)
	}
	res = plusOne([]int{0})
	if !testEq(res, []int{1}) {
		t.Errorf("[1,2,3] plusOne: %v", res)
	}
	res = plusOne([]int{9, 9, 9})
	if !testEq(res, []int{1, 0, 0, 0}) {
		t.Errorf("[1,2,3] plusOne: %v", res)
	}
}

func testEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
