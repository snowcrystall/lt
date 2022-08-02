package exercise

import (
	"testing"
)

func TestNsearch(t *testing.T) {
	a := []int{1, 2, 3}
	r := nsearch(a, 0, 1)
	if r != 0 {
		t.Errorf("%v, find 1 : index: %v", a, r)
	}

	a = []int{1, 2, 3}
	r = nsearch(a, 0, 3)
	if r != 2 {
		t.Errorf("%v, find 1 : index: %v", a, r)
	}

	a = []int{3, 1, 2}
	r = nsearch(a, 1, 1)
	if r != 1 {
		t.Errorf("%v, find 1 : index: %v", a, r)
	}

	a = []int{2, 3, 1}
	r = nsearch(a, 2, 1)
	if r != 2 {
		t.Errorf("%v, find 1 : index: %v", a, r)
	}
}
