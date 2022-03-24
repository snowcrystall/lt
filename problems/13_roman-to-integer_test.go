package problems

import "testing"

func TestRomanToInt(t *testing.T) {
	num := romanToInt("III")
	if num != 3 {
		t.Errorf("III = %d; want 3", num)
	}
	num = romanToInt("IV")
	if num != 4 {
		t.Errorf("IV = %d; want 4", num)
	}
	num = romanToInt("LVIII")
	if num != 58 {
		t.Errorf("LVIII = %d; want 58", num)
	}
	num = romanToInt("MCMXCIV")
	if num != 1994 {
		t.Errorf("LVIII = %d; want 1994", num)
	}
}
