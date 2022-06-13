package simple

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	p := []int{}
	p = append(p, heights...)
	sort.Ints(p)
	fmt.Println(p)
	ret := 0
	for i, v := range p {
		if v != heights[i] {
			ret++
		}
	}
	return ret
}
