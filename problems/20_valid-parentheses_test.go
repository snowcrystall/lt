package problems

import "testing"

func TestIsValid(t *testing.T) {
	s := "()"
	if !isValid(s) {
		t.Errorf("%s isValis false , expect true", s)
	}
	s = "{[]}"
	if !isValid(s) {
		t.Errorf("%s isValis false , expect true", s)
	}
	s = "()[]{}"
	if !isValid(s) {
		t.Errorf("%s isValis false , expect true", s)
	}
	s = "([)]"
	if isValid(s) {
		t.Errorf("%s isValis true , expect false", s)
	}
	s = "("
	if isValid(s) {
		t.Errorf("%s isValis true , expect false", s)
	}
	s = "]"
	if isValid(s) {
		t.Errorf("%s isValis true , expect false", s)
	}
}
