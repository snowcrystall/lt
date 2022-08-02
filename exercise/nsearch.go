package exercise

//原本有序的数组，被循环右移了n位，在这样的数组中查找元素

func nsearch(a []int, nshift int, search int) int {
	n := nshift % len(a)
	r1 := binarySearch(a[:n], search)
	if r1 != -1 {
		return r1
	}
	r2 := binarySearch(a[n:], search)
	if r2 != -1 {
		return r2 + n
	}
	return -1
}

func binarySearch(a []int, r int) int {
	if len(a) == 0 {
		return -1
	}
	low := 0
	high := len(a)

	for low < high {
		m := (low + high) / 2
		if a[m] == r {
			return m
		} else if a[m] < r {
			low = m + 1
		} else {
			high = m
		}
	}
	return -1
}
