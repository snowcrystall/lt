package problems

func plusOne(digits []int) []int {
	tag := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			tag = 0
			break
		} else {
			digits[i] = 0
			tag = 1
		}
	}
	if tag == 1 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
