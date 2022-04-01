package problems

import "testing"

func TestAddBinary(t *testing.T) {
	r := addBinary("101", "1")
	if r != "110" {
		t.Errorf("\"101\" + \"1\" expect 110 , get %v", r)
	}
	r = addBinary("111", "1")
	if r != "1000" {
		t.Errorf("\"101\" + \"1\" expect 1000 , get %v", r)
	}
	r = addBinary("1", "10")
	if r != "11" {
		t.Errorf("\"1\" + \"10\" expect 11 , get %v", r)
	}
	r = addBinary("1", "111")
	if r != "1000" {
		t.Errorf("\"1\" + \"111\" expect 1000 , get %v", r)
	}
}
