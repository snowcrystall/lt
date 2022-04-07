package test

import (
	"fmt"
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestKidsWithCandies(t *testing.T) {
	res := problems.KidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(res)
}
