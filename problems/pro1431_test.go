package problems

import (
	"fmt"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	res := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(res)
}
