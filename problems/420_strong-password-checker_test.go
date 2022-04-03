package problems

import "testing"

func TestStrongPasswordChecker(t *testing.T) {
	str := "a"
	num := strongPasswordChecker(str)
	if 5 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "aA1"
	num = strongPasswordChecker(str)
	if 3 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "1337C0d3"
	num = strongPasswordChecker(str)
	if 0 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "aaa123"
	num = strongPasswordChecker(str)
	if 1 != num {
		t.Errorf("error for  %v, %v", str, num)
	}
	str = "aaaB1"
	num = strongPasswordChecker(str)
	if 1 != num {
		t.Errorf("error for  %v, %v", str, num)
	}
	str = "aaa111"
	num = strongPasswordChecker(str)
	if 2 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "aaa111333"
	num = strongPasswordChecker(str)
	if 3 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "1111111111"
	num = strongPasswordChecker(str)
	if 3 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "ABABABABABABABABABAB1"
	num = strongPasswordChecker(str)
	if 2 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "FFFFFFFFFFFFFFF11111111111111111111AAA"
	num = strongPasswordChecker(str)
	if 23 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "bbaaaaaaaaaaaaaaacccccc"
	num = strongPasswordChecker(str)
	if 8 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "aaaabbbbccccddeeddeeddeedd"
	num = strongPasswordChecker(str)
	if 8 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "aaaaabbbb1234567890ABA"
	num = strongPasswordChecker(str)
	if 3 != num {
		t.Errorf("error for  %v, %v", str, num)
	}

	str = "A1234567890aaabbbbccccc"
	num = strongPasswordChecker(str)
	if 4 != num {
		t.Errorf("error for  %v, %v", str, num)
	}
}
