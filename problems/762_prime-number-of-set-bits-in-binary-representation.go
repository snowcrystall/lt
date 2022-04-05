package problems

import (
	"math"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
	res := 0
	for x := left; x <= right; x++ {
		if isPrime(bits.OnesCount(uint(x))) {
			res++
		}
	}
	return res
}
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	// 只有6x-1和6x+1的数才有可能是质数
	if n%6 != 1 && n%6 != 5 {
		return false
	}
	// 判断这些数能否被小于sqrt(n)的奇数整除
	sqrt := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrt; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
