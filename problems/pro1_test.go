package problems

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 9
	p := twoSum(nums, target)
	fmt.Println(p)

}
