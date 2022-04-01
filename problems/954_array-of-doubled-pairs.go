package problems

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	sort.Ints(arr)
	arrmap := make(map[int]int)
	for _, v := range arr {
		arrmap[v]++
	}

	double := 0

	for _, v := range arr {
		_, ok := arrmap[v]
		if !ok {
			continue
		}

		if arrmap[v] == 1 {
			delete(arrmap, v)
		} else {
			arrmap[v]--
		}
		if v < 0 {
			if v%2 != 0 {
				return false
			}
			double = v / 2
		} else {
			double = v * 2
		}

		v2, ok := arrmap[double]

		if !ok {
			return false
		} else {
			if v2 == 1 {
				delete(arrmap, double)
			} else {
				arrmap[double]--
			}
		}
	}

	return len(arrmap) == 0
}
