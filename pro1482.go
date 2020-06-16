package main

func check(bloomDay []int, m int, k int, day int) bool {
	n := len(bloomDay)
	for j := 0; j < n; j++ {
		ksum := 0
		for d := j; d < j+k && j+k <= n; d++ {
			if bloomDay[d]-day > 0 {
				break
			}
			ksum++
		}
		if ksum == k {
			m--
			j += k - 1
		}
		if m <= 0 {
			return true
		}
	}
	return false
}
func minDays(bloomDay []int, m int, k int) int {
	l := 0
	r := 1000000000
	res := -1
	for l <= r {
		mid := (l + r) >> 1
		if check(bloomDay, m, k, mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
