package problems

//357. 统计各位数字都不同的数字个数
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	nn := 9
	for i := 0; i < n-1; i++ {
		nn *= (9 - i)
	}
	// fmt.Println(nn)
	return nn + countNumbersWithUniqueDigits(n-1)
}
