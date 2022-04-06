package test

import (
	"fmt"
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	problems.Merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)

}
