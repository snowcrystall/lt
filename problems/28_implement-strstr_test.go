package problems

import "testing"

func TestStrStr(t *testing.T) {
	res := strStr("12345", "12")
	if res != 0 {
		t.Errorf("Error find 12345, 12 :%v", res)
	}
	res = strStr("12345", "312")
	if res != -1 {
		t.Errorf("Error find 12345, 312 :%v", res)
	}
	res = strStr("12345", "3")
	if res != 2 {
		t.Errorf("Error find 12345, 3 :%v", res)
	}
	res = strStr("aaa", "aaaa")
	if res != -1 {
		t.Errorf("Error find aaa, aaaa :%v", res)
	}
	res = strStr("a", "a")
	if res != 0 {
		t.Errorf("Error find a, a :%v", res)
	}
}
